# ---- 1. 构建阶段 ----
FROM node:18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制所有源代码
COPY . .

# 构建应用
RUN npm run build

# ---- 2. 运行阶段 ----
FROM nginx:stable-alpine

# 复制构建好的静态文件
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制自定义的 Nginx 配置文件
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80 