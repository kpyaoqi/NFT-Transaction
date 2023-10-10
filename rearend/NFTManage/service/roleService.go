package service

import (
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
)

type RoleService struct {
}

func (roleService RoleService) FindRoleList(page common.Page) []model.TRole {
	//data := page.Data
	//tRole := data.(model.TRole)
	////if data != nil {
	////	if tRole.Active {
	////	}
	////}
	var tRoles []model.TRole
	models.DB.Find(&tRoles)
	return tRoles
}
