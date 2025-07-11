# Mollia API 定义

## 统一约定

-   **Base URL**: `/api/v1`
-   **成功响应**:
    ```json
    {
        "code": 200,
        "msg": "Success",
        "data": { ... }
    }
    ```
-   **失败响应**:
    ```json
    {
        "code": 500, // 或其他错误码
        "msg": "Error message",
        "data": null
    }
    ```
-   **认证**: 需要认证的接口将在 Header 中携带 `Authorization: Bearer <token>`。

---

## 博客展示端 (`mollia-front`)

### 文章 (Article)

-   **获取文章列表（分页）**
    -   **Endpoint**: `GET /articles`
    -   **Query Params**: `page=1&size=10`
    -   **Response**:
        ```json
        {
          "list": [
            {
              "id": 1,
              "title": "文章标题",
              "summary": "文章摘要",
              "author": "作者",
              "createdAt": "2023-01-01T12:00:00Z",
              "tags": ["Go", "Vue"]
            }
          ],
          "total": 100
        }
        ```

-   **根据 ID 获取单篇文章**
    -   **Endpoint**: `GET /articles/:id`
    -   **Response**:
        ```json
        {
          "id": 1,
          "title": "文章标题",
          "content": "<h1>文章内容</h1><p>...</p>",
          "author": "作者",
          "createdAt": "2023-01-01T12:00:00Z",
          "updatedAt": "2023-01-02T12:00:00Z",
          "tags": ["Go", "Vue"],
          "category": "技术"
        }
        ```

### 标签 (Tag)

-   **获取所有标签**
    -   **Endpoint**: `GET /tags`
    -   **Response**: `["Go", "Vue", "Web", ...]`

-   **根据标签获取文章列表**
    -   **Endpoint**: `GET /tags/:tagName/articles`
    -   **Query Params**: `page=1&size=10`
    -   **Response**: 同 `GET /articles`

### 评论 (Comment)

-   **获取文章的评论**
    -   **Endpoint**: `GET /articles/:id/comments`
    -   **Response**:
        ```json
        [
          {
            "id": 1,
            "author": "评论者",
            "content": "这是一条评论",
            "createdAt": "2023-01-01T13:00:00Z"
          }
        ]
        ```

-   **发表评论**
    -   **Endpoint**: `POST /articles/:id/comments`
    -   **Request Body**:
        ```json
        {
          "author": "评论者",
          "content": "这是一条新评论"
        }
        ```

### 个人信息 (Profile)

-   **获取博主信息**
    -   **Endpoint**: `GET /profile`
    -   **Response**:
        ```json
        {
          "name": "博主姓名",
          "avatar": "url/to/avatar.jpg",
          "bio": "个人简介"
        }
        ```
---

## 后台管理端 (`mollia-console`) [需认证]

### 认证 (Auth)

-   **登录**
    -   **Endpoint**: `POST /auth/login`
    -   **Request Body**:
        ```json
        {
          "username": "admin",
          "password": "password"
        }
        ```
    -   **Response**:
        ```json
        {
          "token": "your-jwt-token"
        }
        ```

### 文章 (Article)

-   **获取文章列表（全量，分页）**
    -   **Endpoint**: `GET /admin/articles`
    -   **Query Params**: `page=1&size=10`
    -   **Response**:
        ```json
        {
          "list": [
            {
              "id": 1,
              "title": "文章标题",
              "status": "published", // or "draft"
              "createdAt": "2023-01-01T12:00:00Z",
              "updatedAt": "2023-01-02T12:00:00Z"
            }
          ],
          "total": 100
        }
        ```

-   **创建新文章**
    -   **Endpoint**: `POST /admin/articles`
    -   **Request Body**:
        ```json
        {
          "title": "新文章",
          "content": "...",
          "tags": ["tag1", "tag2"],
          "categoryId": 1,
          "status": "draft"
        }
        ```

-   **根据 ID 获取文章详情**
    -   **Endpoint**: `GET /admin/articles/:id`
    -   **Response**: 同 `GET /articles/:id`

-   **更新文章**
    -   **Endpoint**: `PUT /admin/articles/:id`
    -   **Request Body**: 同 `POST /admin/articles`

-   **删除文章**
    -   **Endpoint**: `DELETE /admin/articles/:id`

### 分类 (Category)

-   **获取所有分类**
    -   **Endpoint**: `GET /admin/categories`

-   **创建分类**
    -   **Endpoint**: `POST /admin/categories`
    -   **Request Body**: `{ "name": "新分类" }`

-   **更新分类**
    -   **Endpoint**: `PUT /admin/categories/:id`
    -   **Request Body**: `{ "name": "更新后的分类" }`

-   **删除分类**
    -   **Endpoint**: `DELETE /admin/categories/:id`


### 存储 (Storage)

-   **上传文件/图片**
    -   **Endpoint**: `POST /admin/storage/upload`
    -   **Request**: `multipart/form-data`
    -   **Response**: `{ "url": "url/to/uploaded/file" }` 