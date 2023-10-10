package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/service"
	"yqnft/NFTManage/utils"
)

type SystemMenuController struct {
}

var menuService service.MenuService
var userRoleMenuService service.UserRoleMenuService

func (systemMenuController SystemMenuController) AllListJson(c *gin.Context) {
	var page common.Page
	c.ShouldBindJSON(&page)
	tMenus := menuService.FindAllMenuListByQueryBean((page.Data), common.Page{})
	menuVOS := utils.ConvertToMenuVO(tMenus)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": menuVOS, "message": "操作成功", "statusCode": 200})
}

func (systemMenuController SystemMenuController) List(c *gin.Context) {
	tMenus := userRoleMenuService.FindAllMenuTree()
	menuVOS := utils.ConvertToMenuVO(tMenus)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": menuVOS, "message": "操作成功", "statusCode": 200})
}
