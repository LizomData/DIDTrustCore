package userController

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	group := e.Group("/api/v1/account")
	group.Use()
	{
		group.POST("/login", loginHandler)
		group.POST("/register", registerHandler)
	}

}
