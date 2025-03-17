package didController

import (
	"DIDTrustCore/did_connnection"
	"DIDTrustCore/model"
	"DIDTrustCore/model/requestBase"
	"github.com/gin-gonic/gin"
)

// CreateSoftwareIdentityApi @Summary 颁发did标识
// @Accept       json
// @Produce      json
// @Param   body body string true "软件名称"
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
	c.JSON(requestBase.ResponseBody(200, "创建成功", "DidID:"+didID))
	return
}

// UpdateSoftwareIdentityApi @Summary 更新did文档
// @Accept       json
// @Produce      json
// @Param   body body model.DidDocument true "新的did文档"
// @Success 200 {object} requestBase.ResponseBodyData "成功"
// @Router /api/v1/did/update_identity [post]
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
// @Param did query string true "didID"
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
		identity, err := QuerySoftwareIdentityByID(did_connnection.FabricClient, did)
		if err != nil {
			c.JSON(requestBase.ResponseBody(500, err.Error(), nil))
			return
		}
		c.JSON(requestBase.ResponseBody(200, "查询成功", identity))
		return
	}
}
