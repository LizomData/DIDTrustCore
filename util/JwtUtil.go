package util

import (
	"DIDTrustCore/model/requestBase"
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

func GenerateToken(userId uint, userName string, expiredTime time.Duration) (string, error) {
	claims := Claims{
		UserID:   userId,
		Username: userName,
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

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
