package userController

import (
	verificationController2 "DIDTrustCore/controller/riskController"
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	var user model.User

	//校验参数
	if err := c.ShouldBindBodyWithJSON(&user); err != nil || !verificationController2.VerifyHeaders(c) || !verifyLoginJson(user) {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "参数错误", gin.H{}))
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

func verifyLoginJson(user model.User) bool {
	return !(user.Username == "" || user.Password == "")
}
