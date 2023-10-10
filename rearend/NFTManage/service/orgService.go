package service

import (
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
)

type OrgService struct {
}

func (orgService OrgService) FindOrgById(orgID string) model.TOrg {
	var org model.TOrg
	models.DB.Where("id=?", orgID).First(&org)
	return org
}

func (orgService OrgService) FindRoleList(org model.TOrg, page common.Page) []model.TOrg {
	userId := "u_10001"
	//userId := config.SessionGet("userId").(string)
	var torgs []model.TOrg
	tOrg := userRoleOrgService.FindOrgByUserId(userId)
	models.DB.Where("comcode like ?", "%"+tOrg[0].ID+"%").Find(&torgs)
	return torgs
}
