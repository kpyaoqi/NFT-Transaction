package service

import (
	"time"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/utils"
)

type UserRoleMenuService struct {
}

var menuService MenuService

func (userRoleMenuService UserRoleMenuService) FindMenuByUserId(userId string) []model.TMenu {
	tRoles := []model.TRole{}
	tMenus := []model.TMenu{}
	models.DB.Table("t_role r").
		Select("r.*").
		Joins("JOIN t_user_role re ON re.roleId = r.id").
		Where("re.userId = ?", userId).
		Where("r.active = ?", 1).
		Order("r.privateOrg DESC, r.sortno DESC").
		Find(&tRoles)
	isExit := make(map[string]bool)
	for _, role := range tRoles {
		tMenus_ := []model.TMenu{}
		models.DB.Table("t_menu m").
			Select("m.*").
			Joins("JOIN t_role_menu re ON re.menuId = m.id").
			Where("re.roleId = ? AND m.active = 1", role.ID).
			Order("m.sortno DESC").
			Find(&tMenus_)
		for _, menu := range tMenus_ {
			if _, ok := isExit[menu.ID]; !ok {
				isExit[menu.ID] = true
				tMenus = append(tMenus, menu)
			}
		}
	}
	return tMenus
}

func (userRoleMenuService UserRoleMenuService) UpdateUserRoles(userId string, rolesString []string) string {
	UserId := config.SessionGet("userId").(string)
	listOld := []string{}
	models.DB.Table("t_user_role").Select("roleId").Where("userId=?", userId).Scan(&listOld)
	models.DB.Where("userId=?", userId).Delete(&model.TUserRole{})
	if len(rolesString) == 0 {
		return ""
	}
	list := []model.TUserRole{}
	for _, roleId := range rolesString {
		ur := model.TUserRole{
			ID:           utils.GetUUID(),
			UserID:       userId, // 假设 userId 已经定义
			RoleID:       roleId, // 假设 roleId 已经定义
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
			CreateUserID: UserId, // 假设 SessionUser 已经定义
			UpdateUserID: UserId, // 假设 SessionUser 已经定义
		}
		list = append(list, ur)
	}
	models.DB.Save(&list)
	return ""
}

func (userRoleMenuService UserRoleMenuService) FindRoleByUserId(userId string) []model.TRole {
	roles := []model.TRole{}
	models.DB.Table("t_role r").Select("r.*").
		Joins("JOIN t_user_role re ON re.roleId=r.id").
		Where("re.userId=? AND r.active=1", userId).
		Order("r.privateOrg DESC,r.sortno DESC").
		Find(&roles)
	return roles
}

func (userRoleMenuService UserRoleMenuService) FindAllMenuTree() []*model.TMenu {
	var menus []*model.TMenu
	models.DB.Where("active=?", 1).Order("sortno DESC").Find(&menus)
	if len(menus) != 0 {
		menuService.MenuList2Tree(menus)
	}
	return menus
}

func (userRoleMenuService UserRoleMenuService) FindMenuByRoleId(roleId string) []*apiResult.MenuVO {
	var menus []*model.TMenu
	models.DB.Table("t_menu m").Select("m.*").
		Joins("JOIN t_role_menu re ON re.menuId=m.id").
		Where("re.roleId=? AND m.active=1", roleId).
		Order("m.sortno DESC").
		Find(&menus)
	menuVOS := utils.ConvertToMenuVO(menus)
	return menuVOS
}
