package service

import (
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	group := e.Group("/api/v1/service")
	group.Use(util.AuthMiddleware())
	{
		group.POST("/checkSoft", checkSoftHandler)
	}
	group.Use(util.AuthMiddlewareV2())
	{
		group.POST("/checkSoftV2", checkSoftHandlerV2)

	}
}

func checkSoftHandler(c *gin.Context) {
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}

func checkSoftHandlerV2(c *gin.Context) {
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}
