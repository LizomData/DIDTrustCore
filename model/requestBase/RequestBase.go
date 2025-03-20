package requestBase

import (
	"github.com/gin-gonic/gin"
)

type ResponseBodyData struct {
	Code int            `json:"code" example:0`   //返回的状态码
	Msg  string         `json:"msg" example:"成功"` //返回的具体信息
	Data map[string]any `json:"data" example:{}`  //返回数据
}

func ResponseBodyBase() (int, gin.H) {
	return 200, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{},
	}
}
func ResponseBody(code int, msg string, data any) (int, gin.H) {
	return 200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
func ResponseBodySuccess(data any) (int, gin.H) {
	return 200, gin.H{
		"code": Success,
		"msg":  "成功",
		"data": data,
	}
}
