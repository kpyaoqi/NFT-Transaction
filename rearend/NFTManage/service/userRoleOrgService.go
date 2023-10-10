package service

import (
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
)

type UserRoleOrgService struct {
}

var orgService OrgService

func (userRoleOrgService UserRoleOrgService) FindOrgIdByUserId(userId string) string {
	tUserOrg := model.TUserOrg{}
	models.DB.Order("managerType DESC").
		Where("userId = ?", userId).
		First(&tUserOrg)
	return tUserOrg.UserID
}

func (userRoleOrgService UserRoleOrgService) UpdateUserOrg(userOrg model.TUserOrg) string {
	userOrgList := userOrg.UserOrgList
	if userOrg.ID == "" || len(userOrgList) == 0 {
		return "数据不能为空"
	}
	models.DB.Where("userId=?", userOrg.UserID).Delete(&model.TUserOrg{})
	models.DB.Save(&userOrgList)
	return ""
}

func (userRoleOrgService UserRoleOrgService) FindOrgByUserId(userId string) []model.TOrg {
	orgs := []model.TOrg{}
	userOrgs := userRoleOrgService.FindUserOrgByUserId(userId)
	for _, userOrg := range userOrgs {
		var org model.TOrg
		org = orgService.FindOrgById(userOrg.OrgID)
		org.ManagerType = userOrg.ManagerType
		orgs = append(orgs, org)
	}
	return orgs
}

func (userRoleOrgService UserRoleOrgService) FindUserOrgByUserId(userId string) []model.TUserOrg {
	var lists []model.TUserOrg
	models.DB.Where("userId=?", userId).
		Order("managerType DESC").
		Find(&lists)
	return lists
}
