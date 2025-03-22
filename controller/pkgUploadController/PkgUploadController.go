package pkgUploadController

import (
	"DIDTrustCore/util"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {

	// 暴露静态文件目录
	e.Static(Uploader.Config.PublicPath, Uploader.Config.UploadDir)

	group := e.Group("/api/v1/pkg")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/upload", upload)
	}
}
