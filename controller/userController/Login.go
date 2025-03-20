package userController

import (
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登陆
// @Accept       json
// @Produce      json
// @Tags 用户管理
// @Param  body body model.User true "登录凭证"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/account/login [post]
func loginHandler(c *gin.Context) {
	vail, user := validateForm(c)
	if !vail {
		return
	}
	isFound, _user := dataBase.FindUser(user.Username)
	//密码校验
	if !isFound || _user.Password != user.Password {
		c.JSON(requestBase.ResponseBody(requestBase.LoginFailed, "用户名或密码错误", gin.H{}))
		return
	}

	user = _user

	// 生成JWT
	tokenString, err := util.GenerateToken(user, 48)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.TokenGenerationFailed, "生成token失败", gin.H{}))
		return
	}

	c.JSON(requestBase.ResponseBodySuccess(gin.H{
		"userInfo": gin.H{"username": user.Username, "privilegeLevel": user.PrivilegeLevel},
		"token":    tokenString,
	}))

}
