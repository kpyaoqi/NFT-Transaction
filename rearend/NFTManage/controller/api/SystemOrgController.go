package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/service"
)

type SystemOrgController struct {
}

var orgService service.OrgService

func (systemOrgController SystemOrgController) Treeselect(c *gin.Context) {
	tOrgs := []model.TOrg{}
	models.DB.Where("pid = ?", "").First(&tOrgs)
	tOrgch := []model.TOrg{}
	models.DB.Where("pid = ?", tOrgs[0].ID).Find(&tOrgch)
	for _, orgch := range tOrgch {
		tOrgs[0].Children = append(tOrgs[0].Children, orgch)
	}
	orgTreeVOS := apiResult.OrgTreeVO{}.OrgTreeConvertOrgTreeVO(tOrgs)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": orgTreeVOS, "message": "操作成功", "statusCode": 200})
}

func (systemOrgController SystemOrgController) PageList(c *gin.Context) {
	var page common.Page
	var org model.TOrg
	c.ShouldBindJSON(&page)
	if page.Data != nil {
		org = (page.Data).(model.TOrg)
	}
	res := orgService.FindRoleList(org, page)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": res, "message": "操作成功", "statusCode": 200})
}

func (systemOrgController SystemOrgController) OrgTypeList(c *gin.Context) {

}
