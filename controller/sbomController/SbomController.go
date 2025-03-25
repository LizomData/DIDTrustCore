package sbomController

import (
	"DIDTrustCore/util"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	// 暴露静态文件目录
	e.Static(Generator.Config.PublicPath, Generator.Config.SBOMStorageDir)

	group := e.Group("/api/v1/sbom")

	group.Use(util.AuthMiddleware())
	{
		group.POST("/generate", generate)

	}

}
