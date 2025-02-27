package common

import (
	"DIDTrustCore/model"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) model.User {
	_user, _ := c.Get("user")
	return _user.(model.User)
}
