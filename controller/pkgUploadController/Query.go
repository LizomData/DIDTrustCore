package pkgUploadController

import (
	"DIDTrustCore/common"
	"DIDTrustCore/model/requestBase"
	pkgDB "DIDTrustCore/util/dataBase/pkgDb"
	"github.com/gin-gonic/gin"
)

// @Summary 查询软件包列表
// @Description 查询用户上传软件包记录
// @Tags 软件包管理
// @Accept json
// @Produce json
// @Param Authorization	header		string	true	"jwt"
// @Param body body QueryRequest true "查询参数"
// @Success 200 {object} model.PkgRecord "软件包列表"
// @Router /api/v1/pkg/query [post]
func query(c *gin.Context) {
	var req QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "无效的请求参数,检查软件包路径", gin.H{}))
		return
	}
	user, _ := common.GetUserFromContext(c)
	records, err := pkgDB.Svc.ListPkgs(user.ID, req.Page, req.Size)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.NotFoundReport, "查询失败:"+err.Error(), gin.H{}))
		return
	}
	c.JSON(requestBase.ResponseBodySuccess(records))
}
