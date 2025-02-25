package userController

import (
	verificationController2 "DIDTrustCore/controller/riskController"
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
	"time"
)

func registerHandler(c *gin.Context) {
	var user model.User

	//参数校验
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

	//非法字符
	if !validateUsername(user.Username) || !validatePassword(user.Password) {
		c.JSON(requestBase.ResponseBody(requestBase.IllegalCharacter, "非法字符", gin.H{}))
		return
	}

	//查询重复
	if isFound, _ := dataBase.FindUser(user.Username); isFound {
		c.JSON(requestBase.ResponseBody(requestBase.RegisterAlready, "用户已被注册", gin.H{}))
		return
	}

	//创建用户
	if err := dataBase.CreateUser(user); err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.RegisterFailed, "注册失败", gin.H{}))
		return
	}

	c.JSON(requestBase.ResponseBodySuccess(gin.H{"username": user.Username, "timeStamp": time.Now().Unix()}))
}
