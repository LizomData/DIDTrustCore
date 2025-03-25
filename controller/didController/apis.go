package didController

import (
	"DIDTrustCore/did_connnection"
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	pkgDB "DIDTrustCore/util/dataBase/pkgDb"
	"github.com/gin-gonic/gin"
)

// CreateSoftwareIdentityApi @Summary 颁发did标识
// @Accept       json
// @Produce      json
// @Tags         did身份标识管理
// @Param   name body string true "软件名称"
// @Description 通过传入软件包名称后台自动生成一个did标识符，并生成与其对应的did文档
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/did/create_identity [post]
func CreateSoftwareIdentityApi(c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(requestBase.ResponseBody(
			400,
			"Invalid JSON",
			nil,
		))
		return
	}
	name, exit := jsonData["name"].(string)
	if !exit {
		c.JSON(requestBase.ResponseBody(400, "name参数残缺", nil))
		return
	}
	doc := CreateDIDDocument(name)
	didID, err := CreateSoftwareIdentity(did_connnection.FabricClient, doc)
	if err != nil {
		c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
		return
	}
	err = pkgDB.Svc.UpdateRecordDidID(name, didID)
	if err != nil {
		c.JSON(requestBase.ResponseBody(500, err.Error(), gin.H{}))
		return
	}

	c.JSON(requestBase.ResponseBodySuccess(gin.H{
		"DidID": didID,
	}))
	return
}

// UpdateSoftwareIdentityApi @Summary 更新did文档
// @Accept       json
// @Produce      json
// @Tags         did身份标识管理
// @Description 通过传入结构体body数据来更新标识信息，注意这个结构体仅仅只是Document这一部分哦，这只是整个标识数据模型中的一部分，切记不要传入整个标识！！！
// @Param   body body model.DidDocument true "新的did文档"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/did/update_identity [put]
func UpdateSoftwareIdentityApi(c *gin.Context) {
	var newDocument model.DidDocument
	if err := c.ShouldBindJSON(&newDocument); err != nil {
		c.JSON(requestBase.ResponseBody(
			400,
			"Invalid JSON",
			nil,
		))
		return
	}
	if err := UpdateDocument(did_connnection.FabricClient, newDocument.ID, newDocument); err != nil {
		c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
		return
	}
	c.JSON(requestBase.ResponseBody(200, "更新文档成功", nil))
	return
}

// QuerySoftwareIdentityApi  @Summary 查询did文档
// @Accept       json
// @Produce      json
// @Tags         did身份标识管理
// @Description 通过didID来查询软件标识，传入指定的didID可以查询到该指定的文档，想要获取链上全部标识时，参数填“all”会返回所有标识。
// @Param did query string true "查询单个标识时传入该文档指定的didID，想要获取全部标识时，就直接传入参数“all”。 "
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/did/query_identity [get]
func QuerySoftwareIdentityApi(c *gin.Context) {
	did := c.Query("did")
	if did == "" {
		c.JSON(requestBase.ResponseBody(
			400,
			"Invalid param",
			nil,
		))
		return
	}
	if did != "all" {
		identity, err := QuerySoftwareIdentityByID(did_connnection.FabricClient, did)
		if err != nil {
			c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
			return
		}
		c.JSON(requestBase.ResponseBody(200, "查询成功", identity))
		return
	} else {
		identity, err := QueryAllSoftwareIdentity(did_connnection.FabricClient)
		if err != nil {
			c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
			return
		}
		c.JSON(requestBase.ResponseBody(200, "查询成功", identity))
		return
	}
}

// RemoveSoftwareIdentityApi  @Summary 吊销did文档
// @Accept       multipart/form-data
// @Produce      json
// @Tags         did身份标识管理
// @Description 通过表单数据来删除标识信息，注意是表单数据哦！！！
// @Param   didID formData string true "didID"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/did/remove_identity [delete]
func RemoveSoftwareIdentityApi(c *gin.Context) {
	didID := c.PostForm("didID")
	if didID == "" {
		c.JSON(requestBase.ResponseBody(400, "未获取到要删除的didID", nil))
		return
	}
	err := DeleteSoftwareByID(did_connnection.FabricClient, didID)
	if err != nil {
		c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
		return
	}
	err = pkgDB.Svc.DeletePkgDidID(didID)
	if err != nil {
		c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
		return
	}

	c.JSON(requestBase.ResponseBody(200, "删除成功", nil))
	return
}
