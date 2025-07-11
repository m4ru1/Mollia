# 定义变量
COMPOSE_FILE = docker-compose.yml

.PHONY: help build up down restart logs ps clean

help:
	@echo "Usage: make [command]"
	@echo "Commands:"
	@echo "  build    构建所有服务的 Docker 镜像"
	@echo "  up       在后台启动所有服务"
	@echo "  down     停止并移除所有服务"
	@echo "  restart  重启所有服务"
	@echo "  logs     查看所有服务的日志"
	@echo "  ps       显示服务的容器状态"
	@echo "  clean    清理已停止的容器和悬空镜像"

# 构建所有 Docker 镜像
build:
	@echo "Building Docker images..."
	docker-compose -f $(COMPOSE_FILE) build

# 在后台启动所有服务
up:
	@echo "Starting all services in background..."
	docker-compose -f $(COMPOSE_FILE) up -d

# 停止并移除所有容器、网络
down:
	@echo "Stopping and removing all services..."
	docker-compose -f $(COMPOSE_FILE) down

# 重启服务
restart: down up

# 查看日志
logs:
	@echo "Showing logs..."
	docker-compose -f $(COMPOSE_FILE) logs -f

# 显示容器状态
ps:
	@echo "Service status:"
	docker-compose -f $(COMPOSE_FILE) ps

# 清理环境
clean: down
	@echo "Cleaning up dangling images..."
	docker image prune -f

# 设置默认目标
default: help 