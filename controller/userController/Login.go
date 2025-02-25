package userController

import (
	verificationController2 "DIDTrustCore/controller/riskController"
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	var user model.User

	//校验参数
	if err := c.ShouldBindBodyWithJSON(&user); err != nil || !verificationController2.VerifyHeaders(c) || !verifyLoginJson(user) {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "参数错误", gin.H{}))
		return
	}

	password_decrypt, err := util.DecryptPassword(user.Password)
	//解密密码
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.IllegalCharacter, "格式不正确", gin.H{}))
		return
	}
	user.Password = password_decrypt

	//密码校验
	if isFound, _user := dataBase.FindUser(user.Username); !isFound || _user.Password != user.Password {
		c.JSON(requestBase.ResponseBody(requestBase.LoginFailed, "用户名或密码错误", gin.H{}))
		return
	}

	// 生成JWT
	tokenString, err := util.GenerateToken(1, user.Username, 48)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.TokenGenerationFailed, "生成token失败", gin.H{}))
		return
	}

	c.JSON(requestBase.ResponseBodySuccess(gin.H{
		"username": user.Username,
		"token":    tokenString,
	}))

}
