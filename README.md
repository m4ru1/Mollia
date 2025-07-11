<div align="center">
  <h1>Mollia - 现代化的 Go + Vue 博客系统</h1>
  <p>
    一个功能完备、前后端分离、容器化部署的现代化博客平台。
  </p>
  <p>
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go Version" />
    <img src="https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=for-the-badge&logo=vue.js" alt="Vue Version" />
    <img src="https://img.shields.io/badge/MySQL-8.0-4479A1?style=for-the-badge&logo=mysql" alt="MySQL Version" />
    <img src="https://img.shields.io/badge/Docker-Powered-2496ED?style=for-the-badge&logo=docker" alt="Docker" />
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" alt="License" />
  </p>
</div>

---

## 简介

**Mollia** 是一个从零开始构建的、完全开源的现代化博客系统。它采用前后端分离的架构，后端使用 Go (Gin) 提供高性能的 RESTful API，前端则分别构建了面向公众的博客展示端和功能强大的后台管理端，均基于 Vue 3 + Vite + Element Plus。

整个项目遵循生产环境的严谨标准进行开发，采用了优雅的分层架构、依赖注入、配置管理和容器化部署，使其成为学习和实践现代 Web 开发的绝佳范例。

## ✨ 主要特性

-   **技术栈**:
    -   **后端**: Go 1.21+, Gin, GORM, JWT, Viper
    -   **前端**: Vue 3, Vite, Pinia, Element Plus, Axios, Cherry-Markdown
    -   **数据库**: MySQL 8.0
-   **前后端分离**: 清晰的 API 边界，便于独立开发和部署。
-   **双前端应用**:
    -   `mollia-front`: 优雅、响应式的博客展示端。
    -   `mollia-console`: 功能齐全的后台管理面板。
-   **完备的管理功能**:
    -   文章管理 (CRUD)
    -   分类管理 (CRUD)
    -   标签管理 (CRUD)
    -   安全的 JWT 用户认证。
-   **强大的编辑器**: 集成 `cherry-markdown`，支持 Markdown 实时预览和图片无缝上传。
-   **生产级后端架构**:
    -   **分层架构**: 清晰的 `Handler -> Service -> Repository` 结构。
    -   **依赖注入**: 高度模块化，易于测试和维护。
    -   **优雅停机**: 保证服务在关闭时不会丢失数据。
    -   **配置管理**: 通过环境变量和 `.env` 文件管理配置，安全可靠。
-   **容器化**:
    -   通过多阶段构建优化 `Dockerfile`，生成轻量级镜像。
    -   使用 `docker-compose` 一键编排和部署整个应用栈 (Backend, Frontend, Admin, DB)。
-   **便捷的开发体验**: 提供 `Makefile` 来简化构建、运行和清理等常用操作。

## 🏛️ 系统架构

Mollia 包含四个核心的容器化服务，通过 `docker-compose` 进行编排：

```mermaid
graph TD
    subgraph "用户浏览器"
        U1(访客)
        U2(管理员)
    end

    subgraph "Nginx 容器"
        N1[front (Port 80)]
        N2[console (Port 81)]
    end

    subgraph "Go 后端容器"
        S[server (Port 8080)]
    end

    subgraph "数据库容器"
        DB[(MySQL DB)]
    end

    U1 -- "访问博客" --> N1
    U2 -- "访问后台" --> N2
    N1 -- "API 请求" --> S
    N2 -- "API 请求" --> S
    S <--> DB
```

-   **mollia-front (localhost:80)**: 博客展示端，由 Nginx 提供服务。
-   **mollia-console (localhost:81)**: 后台管理端，由 Nginx 提供服务。
-   **mollia-server (localhost:8080)**: Go 后端 API 服务。
-   **mollia_db**: MySQL 数据库，数据持久化。

## 🚀 快速开始

### 先决条件

在开始之前，请确保您的系统中已安装以下软件：
-   [Docker](https://www.docker.com/get-started)
-   [Docker Compose](https://docs.docker.com/compose/install/)
-   `make` (在 Linux/macOS 中通常自带，Windows 用户可能需要安装 [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm))

### 安装与运行

1.  **克隆项目**
    ```bash
    git clone https://github.com/m4ru1/Mollia.git
    cd Mollia
    ```

2.  **创建配置文件**
    在项目根目录创建一个名为 `.env` 的文件。您可以复制 `docker-compose.yml` 中的环境变量作为模板，并修改其中的密码和密钥。这是一个必要的安全步骤。
    
    **.env 文件示例:**
    ```env
    # 数据库配置
    MYSQL_ROOT_PASSWORD=your_very_strong_password
    MYSQL_DATABASE=mollia

    # 后端服务配置 (docker-compose 会自动替换 db 为容器名)
    DB_DSN=root:${MYSQL_ROOT_PASSWORD}@tcp(db:3306)/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local
    JWT_SECRET=a_very_secret_and_long_key_for_production
    PORT=8080
    ```

3.  **构建并启动服务**
    使用 `Makefile` 中的命令一键完成所有操作。

    ```bash
    # 构建所有服务的 Docker 镜像
    make build

    # 在后台启动所有容器
    make up
    ```
    服务启动后：
    -   博客前台: **http://localhost**
    -   后台管理: **http://localhost:81**

4.  **默认登录凭据**
    -   **用户名**: `admin`
    -   **密码**: `password`

## ⚙️ Makefile 命令

我们提供了一个方便的 `Makefile` 来管理应用的生命周期：

| 命令 | 描述 |
| :--- | :--- |
| `make help` | 显示所有可用的命令。 |
| `make build` | 构建所有服务的 Docker 镜像。 |
| `make up` | 在后台启动所有服务。 |
| `make down` | 停止并移除所有容器、网络。 |
| `make restart` | 重启所有服务 (`down` 然后 `up`)。 |
| `make logs` | 实时查看所有服务的日志。 |
| `make ps` | 显示所有容器的运行状态。 |
| `make clean` | 停止所有服务并清理悬空镜像。 |


## 📂 项目结构

```
Mollia/
├── docker/                 # Docker 相关配置
│   └── db/init/            # 数据库初始化脚本
├── mollia-console/         # 后台管理端 (Vue)
├── mollia-front/           # 博客展示端 (Vue)
├── mollia-server/          # 后端服务 (Go)
│   ├── api/                # 路由注册
│   ├── internal/           # 内部业务逻辑
│   │   ├── handler/
│   │   ├── middleware/
│   │   ├── model/
│   │   ├── repository/
│   │   └── service/
│   └── pkg/                # 公用库 (配置, 数据库)
├── .env                    # (需要手动创建) 环境变量
├── API.md                  # API 文档
├── docker-compose.yml      # Docker 服务编排
└── Makefile                # 便捷的构建命令
```

## 📜 API 文档

详细的 RESTful API 定义请参考 `API.md` 文件。

## 🤝 贡献

欢迎任何形式的贡献！如果您有好的想法或发现了 bug，请随时提交 Pull Request 或创建 Issue。

## 📄 许可证

本项目基于 [MIT License](LICENSE) 开源。
