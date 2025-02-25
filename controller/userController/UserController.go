package userController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"unicode"
)

func Routers(e *gin.Engine) {
	group := e.Group("/api/v1/account")
	group.Use()
	{
		group.POST("/login", loginHandler)
		group.POST("/register", registerHandler)
	}

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
