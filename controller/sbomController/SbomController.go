package sbomController

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	group := e.Group("/api/v1/sbom")
	group.Use()
	{
		group.POST("/generate", generate)
	}

}
