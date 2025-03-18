package didController

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	group := e.Group("/api/v1/did")
	group.Use()
	{
		group.POST("/create_identity", CreateSoftwareIdentityApi)
		group.GET("/query_identity", QuerySoftwareIdentityApi)
		group.PUT("/update_document", UpdateSoftwareIdentityApi)
		group.PUT("/update_key")
		group.DELETE("/remove_identity", RemoveSoftwareIdentityApi)
	}
}
