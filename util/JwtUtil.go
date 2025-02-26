package util

import (
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT密钥
var jwtSecret = []byte("2004qwe")

// JWT Claims
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(user model.User, expiredTime time.Duration) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * expiredTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(requestBase.ResponseBody(requestBase.LoginStatusInvalid, "登录状态失效", gin.H{}))
			c.Abort()
			return
		}
		//tokenString := authHeader[len("Bearer "):]
		tokenString := authHeader
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(requestBase.ResponseBody(requestBase.InvalidTokens, "无效token", gin.H{}))
			c.Abort()
			return
		}

		isFound, user := dataBase.FindUserById(claims.UserID)
		if !isFound {
			c.JSON(requestBase.ResponseBody(requestBase.NotUser, "用户不存在", gin.H{}))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// JWT中间件V2
func AuthMiddlewareV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(requestBase.ResponseBody(requestBase.LoginStatusInvalid, "登录状态失效", gin.H{}))
			c.Abort()
			return
		}
		//tokenString := authHeader[len("Bearer "):]
		tokenString := authHeader
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(requestBase.ResponseBody(requestBase.InvalidTokens, "无效token", gin.H{}))
			c.Abort()
			return
		}

		isFound, user := dataBase.FindUserById(claims.UserID)
		if !isFound {
			c.JSON(requestBase.ResponseBody(requestBase.NotUser, "用户不存在", gin.H{}))
			c.Abort()
			return
		}

		if user.PrivilegeLevel == 0 {
			c.JSON(requestBase.ResponseBody(requestBase.NotPrivileged, "权限不足", gin.H{}))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
