# ---- 1. 构建阶段 ----
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制所有源代码
COPY . .

# 编译应用
# CGO_ENABLED=0 是为了静态编译，避免在 alpine 环境中缺少 C 库
# -o /server 是将编译好的二进制文件输出到根目录下的 server 文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server .

# ---- 2. 运行阶段 ----
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /server .

# 从构建阶段复制上传目录的占位符（如果需要包含默认上传内容）
# 这里我们创建一个空的，因为上传是运行时行为
RUN mkdir -p /app/uploads

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./server"] 