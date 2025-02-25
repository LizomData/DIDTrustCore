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
		group.POST("/checkSoftV2", checkSoftHandlerV2)
	}
}

func checkSoftHandler(c *gin.Context) {
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}
func checkSoftHandlerV2(c *gin.Context) {
	user := requestBase.GetUserFromContext(c)

	if user.PrivilegeLevel == 0 {
		c.JSON(requestBase.ResponseBody(requestBase.NotPrivileged, "权限不足", gin.H{}))
		return
	}
	c.JSON(requestBase.ResponseBodySuccess(gin.H{"checkStatus": true}))
}
