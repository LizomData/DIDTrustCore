package userController

import (
	"DIDTrustCore/common"
	"DIDTrustCore/model/requestBase"
	"github.com/gin-gonic/gin"
)

// @Summary 用户信息
// @Accept       json
// @Produce      json
// @Tags 用户管理
// @Param Authorization	header		string	true	"jwt"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/account/getUserInfo [get]
func getUserInfo(c *gin.Context) {
	user := common.GetUserFromContext(c)
	c.JSON(requestBase.ResponseBodySuccess(gin.H{
		"userInfo": gin.H{"userId": user.ID, "username": user.Username, "privilegeLevel": user.PrivilegeLevel}}))

}
