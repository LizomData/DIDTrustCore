package routers

import (
	"DIDTrustCore/util"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(util.Cors())
	for _, opt := range options {
		opt(r)
	}
	return r
}
