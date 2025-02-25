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
}

func checkSoftHandler(c *gin.Context) {
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}
