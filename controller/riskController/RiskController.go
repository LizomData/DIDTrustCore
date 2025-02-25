package riskController

import (
	"github.com/gin-gonic/gin"
)

func VerifyHeaders(c *gin.Context) bool {
	return true
	s := c.Request.Header.Get("S")
	t := c.Request.Header.Get("T")
	return !(s == "" || t == "")
}
