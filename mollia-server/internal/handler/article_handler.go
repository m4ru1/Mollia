package handler

import (
	"context"
	"mollia/internal/model"
	"mollia/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticleHandler 封装了文章相关的处理器
type ArticleHandler struct {
	svc service.ArticleService
}

// NewArticleHandler 创建一个新的 ArticleHandler
func NewArticleHandler(svc service.ArticleService) *ArticleHandler {
	return &ArticleHandler{svc: svc}
}

// GetArticleByID 公共接口：根据 ID 获取单篇文章
func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的文章ID"})
		return
	}
	article, err := h.svc.GetArticle(c.Request.Context(), uint(id))
	if err != nil {
		// 假设 service 层会处理好 sql.ErrNoRows，并返回一个可识别的错误类型
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": article})
}

// GetArticles 公共接口：获取文章列表
func (h *ArticleHandler) GetArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	offset := (page - 1) * size

	articles, total, err := h.svc.ListArticles(c.Request.Context(), size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"list": articles, "total": total}})
}

// --- Admin Handlers ---

// CreateArticle 管理接口：创建文章
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var req model.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数无效: " + err.Error()})
		return
	}
	ctx := context.WithValue(c.Request.Context(), "username", c.GetString("username"))
	id, err := h.svc.CreateArticle(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "创建文章失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"id": id}})
}

// UpdateArticle 管理接口：更新文章
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的文章ID"})
		return
	}
	var req model.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数无效: " + err.Error()})
		return
	}

	if err := h.svc.UpdateArticle(c.Request.Context(), uint(id), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "更新文章失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "更新成功"})
}

// DeleteArticle 管理接口：删除文章
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的文章ID"})
		return
	}
	if err := h.svc.DeleteArticle(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "删除文章失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})
}

// GetAdminArticles 管理接口：获取文章列表
func (h *ArticleHandler) GetAdminArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	offset := (page - 1) * size

	articles, total, err := h.svc.ListAdminArticles(c.Request.Context(), size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"list": articles, "total": total}})
}

// GetAdminArticleByID 管理接口：根据ID获取单篇文章
func (h *ArticleHandler) GetAdminArticleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的文章ID"})
		return
	}
	article, err := h.svc.GetAdminArticleByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": article})
}
