package userController

import (
	"DIDTrustCore/util"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	group := e.Group("/api/v1/account")
	group.Use()
	{
		group.POST("/login", loginHandler)
		group.POST("/register", registerHandler)
	}
	group.Use(util.AuthMiddleware())
	{
		group.GET("/getUserInfo", getUserInfo)
	}

}
