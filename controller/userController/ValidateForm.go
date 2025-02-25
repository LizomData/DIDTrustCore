package userController

import (
	verificationController2 "DIDTrustCore/controller/riskController"
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"unicode"
)

func validateForm(c *gin.Context) (bool, model.User) {

	var user model.User

	//校验参数
	if err := c.ShouldBindBodyWithJSON(&user); err != nil || !verificationController2.VerifyHeaders(c) {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "参数错误", gin.H{}))
		return false, user
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "用户名或密码不能为空", gin.H{}))
		return false, user
	}

	//解密密码
	if _, err := DecryptPassword(&user); err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.IllegalCharacter, "格式不正确", gin.H{}))
		return false, user
	}

	//非法字符
	if !validateUsername(user.Username) || !validatePassword(user.Password) {
		c.JSON(requestBase.ResponseBody(requestBase.IllegalCharacter, "非法字符", gin.H{}))
		return false, user
	}
	return true, user
}

func validateUsername(username string) bool {
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
func validatePassword(password string) bool {
	// 定义允许的字符集：字母、数字和一些特殊字符
	allowedPattern := `^[a-zA-Z0-9!@#$%^&*().]+$`
	matched, err := regexp.MatchString(allowedPattern, password)
	if err != nil {
		fmt.Println("正则表达式错误:", err)
		return false
	}
	return matched

}
