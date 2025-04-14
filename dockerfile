# 阶段1：编译 Go 二进制
FROM golang:1.24-bullseye AS builder
WORKDIR /app

# RUN apk --no-cache --update add ca-certificates

COPY go.mod go.sum ./
RUN go mod download

# 将 org1.example.com 复制到构建环境的 /app 目录下
COPY . .
COPY org1.example.com /app/org1.example.com/

# 编译时将二进制输出到 /app 目录
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o /app/gin-app

# 阶段2：运行镜像
FROM alpine:3.18

# 创建应用目录并保持与构建环境相同的路径
WORKDIR /app

# 完整复制构建环境的 /app 目录结构
COPY --from=builder /app .


EXPOSE 8000

# 确保以正确权限运行
CMD ["/app/gin-app"]