package repository

import (
	"context"
	"database/sql"
	"mollia/internal/model"
	"strings"
)

type mysqlArticleRepository struct {
	db *sql.DB
}

// NewMySQLArticleRepository 创建一个新的 ArticleRepository 实现
func NewMySQLArticleRepository(db *sql.DB) ArticleRepository {
	return &mysqlArticleRepository{db: db}
}

func (r *mysqlArticleRepository) GetByID(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	var article model.ArticleResponse
	query := `
		SELECT a.id, a.title, a.content, a.created_at, a.updated_at, u.username, cat.name
		FROM articles a
		JOIN users u ON a.author_id = u.id
		JOIN categories cat ON a.category_id = cat.id
		WHERE a.id = ? AND a.status = 'published'
	`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&article.ID, &article.Title, &article.Content,
		&article.CreatedAt, &article.UpdatedAt, &article.Author, &article.Category,
	)
	if err != nil {
		return nil, err // 直接返回错误，让上层处理
	}

	tagsQuery := `
		SELECT t.name FROM tags t
		JOIN article_tags at ON t.id = at.tag_id
		WHERE at.article_id = ?
	`
	rows, err := r.db.QueryContext(ctx, tagsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tagName string
		if err := rows.Scan(&tagName); err != nil {
			return nil, err
		}
		tags = append(tags, tagName)
	}
	article.Tags = tags

	return &article, nil
}

// NOTE: 下面的实现为了演示方便，省略了部分细节，如事务处理和完整的错误处理

func (r *mysqlArticleRepository) List(ctx context.Context, limit, offset int) ([]model.ArticleInList, int64, error) {
	var total int64
	countQuery := "SELECT COUNT(*) FROM articles WHERE status = 'published'"
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT a.id, a.title, a.summary, a.created_at, u.username
		FROM articles a
		JOIN users u ON a.author_id = u.id
		WHERE a.status = 'published'
		ORDER BY a.created_at DESC
		LIMIT ? OFFSET ?
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	articles := make([]model.ArticleInList, 0)
	for rows.Next() {
		var article model.ArticleInList
		if err := rows.Scan(&article.ID, &article.Title, &article.Summary, &article.CreatedAt, &article.Author); err != nil {
			return nil, 0, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	if len(articles) > 0 {
		if err := r.fetchTagsForArticleList(ctx, articles); err != nil {
			return nil, 0, err
		}
	}

	return articles, total, nil
}

func (r *mysqlArticleRepository) Create(ctx context.Context, req *model.CreateArticleRequest, authorID uint) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// defer an automatic rollback in case of error
	defer tx.Rollback()

	// 插入文章
	articleQuery := "INSERT INTO articles (title, content, summary, author_id, category_id, status) VALUES (?, ?, ?, ?, ?, ?)"
	summary := req.Content
	if len(summary) > 200 {
		summary = string([]rune(summary)[:200])
	}
	res, err := tx.ExecContext(ctx, articleQuery, req.Title, req.Content, summary, authorID, req.CategoryID, req.Status)
	if err != nil {
		return 0, err
	}
	articleID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 处理标签
	for _, tagName := range req.Tags {
		var tagID int64
		err := tx.QueryRowContext(ctx, "SELECT id FROM tags WHERE name = ?", tagName).Scan(&tagID)
		if err == sql.ErrNoRows {
			tagRes, err := tx.ExecContext(ctx, "INSERT INTO tags (name) VALUES (?)", tagName)
			if err != nil {
				return 0, err
			}
			tagID, _ = tagRes.LastInsertId()
		} else if err != nil {
			return 0, err
		}

		// 关联文章和标签
		_, err = tx.ExecContext(ctx, "INSERT INTO article_tags (article_id, tag_id) VALUES (?, ?)", articleID, tagID)
		if err != nil {
			return 0, err
		}
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return articleID, nil
}

func (r *mysqlArticleRepository) Update(ctx context.Context, id uint, req *model.UpdateArticleRequest) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 更新文章表
	articleQuery := "UPDATE articles SET title = ?, content = ?, summary = ?, category_id = ?, status = ? WHERE id = ?"
	summary := req.Content
	if len(summary) > 200 {
		summary = string([]rune(summary)[:200])
	}
	_, err = tx.ExecContext(ctx, articleQuery, req.Title, req.Content, summary, req.CategoryID, req.Status, id)
	if err != nil {
		return err
	}

	// 更新标签关联：先删除旧的
	_, err = tx.ExecContext(ctx, "DELETE FROM article_tags WHERE article_id = ?", id)
	if err != nil {
		return err
	}

	// 再插入新的
	for _, tagName := range req.Tags {
		var tagID int64
		err := tx.QueryRowContext(ctx, "SELECT id FROM tags WHERE name = ?", tagName).Scan(&tagID)
		if err == sql.ErrNoRows {
			tagRes, err := tx.ExecContext(ctx, "INSERT INTO tags (name) VALUES (?)", tagName)
			if err != nil {
				return err
			}
			tagID, _ = tagRes.LastInsertId()
		} else if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, "INSERT INTO article_tags (article_id, tag_id) VALUES (?, ?)", id, tagID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *mysqlArticleRepository) Delete(ctx context.Context, id uint) (int64, error) {
	query := "DELETE FROM articles WHERE id = ?"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *mysqlArticleRepository) ListAdmin(ctx context.Context, limit, offset int) ([]model.AdminArticleInList, int64, error) {
	// 逻辑与 List 类似，但查询的字段和表不同，且无 status 过滤
	// 同样需要解决 N+1 问题
	return nil, 0, nil
}

func (r *mysqlArticleRepository) GetAdminByID(ctx context.Context, id uint) (*model.ArticleResponse, error) {
	// 逻辑与 GetByID 类似，但无 status 过滤
	return nil, nil
}

// fetchTagsForArticleList 使用一条 SQL 查询为多篇文章获取标签，以解决 N+1 问题
func (r *mysqlArticleRepository) fetchTagsForArticleList(ctx context.Context, articles []model.ArticleInList) error {
	if len(articles) == 0 {
		return nil
	}

	articleIDs := make([]interface{}, len(articles))
	articleMap := make(map[uint]*model.ArticleInList, len(articles))
	for i, a := range articles {
		articleIDs[i] = a.ID
		// Take a pointer to the slice element
		articles[i].Tags = []string{}
		articleMap[a.ID] = &articles[i]
	}

	tagsQuery := `
		SELECT at.article_id, t.name FROM article_tags at
		JOIN tags t ON at.tag_id = t.id
		WHERE at.article_id IN (?` + strings.Repeat(",?", len(articleIDs)-1) + `)
	`

	rows, err := r.db.QueryContext(ctx, tagsQuery, articleIDs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var articleID uint
		var tagName string
		if err := rows.Scan(&articleID, &tagName); err != nil {
			return err
		}
		if article, ok := articleMap[articleID]; ok {
			article.Tags = append(article.Tags, tagName)
		}
	}
	return rows.Err()
}
