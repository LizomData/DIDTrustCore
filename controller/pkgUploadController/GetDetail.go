package pkgUploadController

import (
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util/dataBase/sbomDb"
	"DIDTrustCore/util/dataBase/scanReportDb"
	"github.com/gin-gonic/gin"
)

// @Summary 获取软件包相关信息(生成的sbom,漏洞分析报告下载路径)
// @Description 获取软件包相关信息(生成的sbom,漏洞分析报告下载路径)
// @Tags 软件包管理
// @Accept json
// @Produce json
// @Param Authorization	header		string	true	"jwt"
// @Param body body GetDetailRequest true "查询参数"
// @Success 200 {object} GetDetailResponse "结果"
// @Router /api/v1/pkg/getDetail [post]
func getDetail(c *gin.Context) {
	var req GetDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "无效的请求参数,检查请求参数", gin.H{}))
		return
	}
	if req.DidID == "" {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "didid为空", gin.H{}))
		return
	}

	sbom_record, err := sbomDb.Sbom_svc.GetSBOMByDidID(req.DidID)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.NotFoundReport, "查询失败:"+err.Error(), gin.H{}))
		return
	}

	scan_record, err := scanReportDb.Svc.GetRecordByDidID(req.DidID)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.NotFoundReport, "查询失败:"+err.Error(), gin.H{}))
		return
	}
	c.JSON(requestBase.ResponseBodySuccess(GetDetailResponse{sbom_record.DownloadURL, scan_record.DownloadURL}))

}
