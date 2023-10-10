package service

import (
	"time"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/utils"
)

type UserService struct {
}

var userRoleOrgService UserRoleOrgService
var userRoleMenuService UserRoleMenuService

func (userService UserService) UserTypeList() []apiResult.LVDTO {
	lvdtos := []apiResult.LVDTO{}
	models.DB.Table("t_dic_data").
		Where("active = ? AND val IS NOT NULL AND typekey = userType", 1).
		Select("name as label, val as value").
		Find(&lvdtos)
	return lvdtos
}

func (userService UserService) SaveUser(user model.TUser) model.TUser {
	if user.ID == "" {
		user.ID = utils.GetTimeNO()
		var count int64
		models.DB.Table("t_user").Where("account=?", user.Account).Count(&count)
		if count != 0 {
			panic("用户账号" + user.Account + "已存在！")
		}
	}
	user.Password = utils.EncoderByMd5With32Bit(user.Password)
	if user.CreateTime.IsZero() {
		user.CreateTime = time.Now()
	}
	if user.Avatar == "" {
		user.Avatar = "/avatar/default.jpg"
	}
	user.UpdateTime = time.Now()
	userId := config.SessionGet("userId").(string)
	user.CreateUserID = userId
	user.UpdateUserID = userId
	user.Address = utils.PhoneToEVMAddress(user.Mobile)
	models.DB.Save(&user)
	if len(user.OrgList) == 0 {
		user.OrgList = append(user.OrgList, model.TOrg{ID: "defaultOrgId"})
	}
	userOrgList := []model.TUserOrg{}
	for _, org := range user.OrgList {
		userOrg := model.TUserOrg{}
		userOrg.ID = utils.GetUUID()
		userOrg.CreateTime = time.Now()
		userOrg.UpdateTime = time.Now()
		userOrg.CreateUserID = userId
		userOrg.UpdateUserID = userId
		userOrg.UserID = user.ID
		userOrg.OrgID = org.ID
		userOrg.ManagerType = org.ManagerType
		userOrgList = append(userOrgList, userOrg)
	}
	saveUserOrg := model.TUserOrg{}
	saveUserOrg.UserID = user.ID
	saveUserOrg.UserOrgList = userOrgList
	userRoleOrgService.UpdateUserOrg(saveUserOrg)
	rolesString := []string{}
	if len(user.Roles) == 0 {
		rolesString = append(rolesString, "nft_role")
		//rolesString = append(rolesString, "defaultRoleId")
	} else {
		for _, role := range user.Roles {
			rolesString = append(rolesString, role.ID)
		}
	}
	userRoleMenuService.UpdateUserRoles(user.ID, rolesString)
	return user
}

func (userService UserService) FollowWorks(worksId string) string {
	tUserFollow := model.TUserFollow{}
	userId := config.SessionGet("userId").(string)

	models.DB.Table("t_user_follow").
		Where("userId=? AND worksId=?", userId, worksId).
		First(&tUserFollow)
	if tUserFollow.ID == "" {
		tUserFollow.ID = utils.GetTimeNO()
		tUserFollow.UserID = userId
		tUserFollow.WorksID = worksId
		tUserFollow.IsValid = 1
		tUserFollow.CreateTime = time.Now()
		tUserFollow.UpdateTime = time.Now()
		models.DB.Save(&tUserFollow)
		return "收藏成功"
	} else if tUserFollow.IsValid == 0 {
		tUserFollow.UpdateTime = time.Now()
		tUserFollow.IsValid = 1
		models.DB.Save(&tUserFollow)
		return "收藏成功"
	} else {
		tUserFollow.UpdateTime = time.Now()
		tUserFollow.IsValid = 0
		models.DB.Save(&tUserFollow)
		return "取消收藏成功"
	}
}

func (userService UserService) HisTranInUser() []apiResult.UserHisTranVO {
	vos := []apiResult.UserHisTranVO{}
	userId := config.SessionGet("userId").(string)
	models.DB.Table("nft_order").
		Where("payState = 1").
		Where("fromUser=? OR toUser = ?", userId, userId).
		Order("txTime DESC").Limit(8).
		Scan(&vos)
	return vos
}

func (userService UserService) FollowStatus(worksId string) int64 {
	userId := config.SessionGet("userId").(string)
	var status int64
	models.DB.Table("t_user_follow").
		Select("isValid").
		Where("worksId=? AND userId=?", worksId, userId).
		Scan(&status)
	return status
}

func (userService UserService) GetUserById(userId string) model.TUser {
	var user model.TUser
	models.DB.Where("id=?", userId).First(&user)
	return user
}

func (userService UserService) FindUserById(userId string) model.TUser {
	user := userService.GetUserById(userId)
	orgByUserId := userRoleOrgService.FindOrgByUserId(userId)
	user.OrgList = orgByUserId
	roles := userRoleMenuService.FindRoleByUserId(userId)
	user.Roles = roles
	return user
}

func (userService UserService) FindUserList() []model.TUser {
	var users []model.TUser
	models.DB.Find(&users)
	for i := range users {
		userId := users[i].ID
		users[i].OrgList = userRoleOrgService.FindOrgByUserId(userId)
		users[i].Roles = userRoleMenuService.FindRoleByUserId(userId)
	}
	return users
}
