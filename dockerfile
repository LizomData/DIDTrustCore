# 阶段1：编译 Go 二进制
FROM golang:1.24-alpine AS builder
WORKDIR /app

# 1. 配置国内镜像代理加速
ENV GOPROXY=https://goproxy.cn,direct

# 2. 安装 CA 证书（解决 HTTPS 验证问题）
RUN apk add --no-cache ca-certificates


# 3. 分离依赖下载与代码构建
COPY go.mod go.sum ./
RUN go mod download


COPY . .
COPY org1.example.com /app/

# 4. 优化构建参数
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -o /gin-app

# 阶段2：最小化运行镜像
FROM alpine:3.18
RUN apk add --no-cache tzdata

# 5. 从 builder 镜像继承 CA 证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /gin-app /gin-app
EXPOSE 8000
CMD ["/gin-app"]