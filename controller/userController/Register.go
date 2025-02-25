package userController

import (
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
	"time"
)

func registerHandler(c *gin.Context) {

	vail, user := validateForm(c)
	if !vail {
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
