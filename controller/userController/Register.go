package userController

import (
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
	"time"
)

// @Summary 用户注册
// @Accept       json
// @Produce      json
// @Tags 用户管理
// @Param  body body model.User true "注册凭证"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/account/register [post]
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

	finded, user_new := dataBase.FindUser(user.Username)
	if !finded {
		c.JSON(requestBase.ResponseBody(requestBase.RegisterFailed, "注册失败", gin.H{}))
		return
	}
	// 生成JWT
	tokenString, err := util.GenerateToken(user_new, 240)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.TokenGenerationFailed, "生成token失败", gin.H{}))
		return
	}

	c.JSON(requestBase.ResponseBodySuccess(gin.H{"username": user.Username, "token": tokenString, "timeStamp": time.Now().Unix()}))
}
