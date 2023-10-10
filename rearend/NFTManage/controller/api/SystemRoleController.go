package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	"yqnft/NFTManage/service"
)

type SystemRoleController struct {
	SystemMenuController
}

var roleService service.RoleService

func (systemRoleController SystemRoleController) List(c *gin.Context) {
	tRoles := []model.TRole{}
	models.DB.Where("active = 1").Find(&tRoles)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": tRoles, "message": "操作成功", "statusCode": 200})
}

func (systemRoleController SystemRoleController) PageList(c *gin.Context) {
	var page common.Page
	c.ShouldBindJSON(&page)
	res := roleService.FindRoleList(page)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": res, "message": "操作成功", "statusCode": 200})
}

func (systemRoleController SystemRoleController) GetMenusByRoleId(c *gin.Context) {
	roleId := c.Query("roleId")
	res := userRoleMenuService.FindMenuByRoleId(roleId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": res, "message": "操作成功", "statusCode": 200})
}
