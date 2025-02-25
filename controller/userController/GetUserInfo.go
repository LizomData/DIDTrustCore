package userController

import (
	"DIDTrustCore/model/requestBase"
	"github.com/gin-gonic/gin"
)

func getUserInfo(c *gin.Context) {
	user := requestBase.GetUserFromContext(c)
	c.JSON(requestBase.ResponseBodySuccess(gin.H{
		"userInfo": gin.H{"userId": user.ID, "username": user.Username, "privilegeLevel": user.PrivilegeLevel}}))

}
