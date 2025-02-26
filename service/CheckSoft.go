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

// @Summary 功能api
// @Accept       json
// @Produce      json
// @Param Authorization	header		string	true	"jwt"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/service/checkSoft [post]
func checkSoftHandler(c *gin.Context) {
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}

// @Summary 功能api2
// @Accept       json
// @Produce      json
// @Param Authorization	header		string	true	"jwt"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/service/checkSoftV2 [post]
func checkSoftHandlerV2(c *gin.Context) {
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}
