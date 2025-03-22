package sbomController

import (
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util/dataBase"
	"github.com/gin-gonic/gin"
)

// @Summary 查询SBOM接口
// @Description 查找用户生成sbom历史记录
// @Tags SBOM管理
// @Accept json
// @Produce json
// @Param Authorization	header		string	true	"jwt"
// @Param body body QuerySBOMRequest true "查询参数"
// @Success 200 {object} model.SBOMReport "SBOM记录"
// @Router /api/v1/sbom/query [post]
func query(c *gin.Context) {
	var req QuerySBOMRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "无效的请求参数,检查查询参数", gin.H{}))
		return
	}
	record, err := dataBase.Sbom_svc.GetSBOMByID(req.SbomReportId)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "查询错误"+err.Error(), gin.H{}))
		return
	}
	c.JSON(requestBase.ResponseBodySuccess(record))
}
