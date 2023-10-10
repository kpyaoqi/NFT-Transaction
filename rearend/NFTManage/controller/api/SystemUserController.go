package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/service"
	"yqnft/NFTManage/utils"
)

type SystemUserController struct {
}

var userService service.UserService

func (con SystemUserController) Menu(c *gin.Context) {
	userId := config.SessionGet("userId").(string)
	tMenus := userRoleMenuService.FindMenuByUserId(userId)
	menuResults := []apiResult.MenuVO{}
	value, _ := c.Get("fromLogin")
	if value == true {
		for _, menu := range tMenus {
			menuResult := apiResult.MenuVO{}
			utils.StructAssign(&menuResult, &menu)
			menuResults = append(menuResults, menuResult)
		}
		c.Set("menuResults", menuResults)
	} else {
		for _, menu := range tMenus {
			if menu.MenuType == 0 {
				continue
			}
			menuResult := apiResult.MenuVO{}
			utils.StructAssign(&menuResult, &menu)
			menuResults = append(menuResults, menuResult)
		}
		c.Set("menuResults", menuResults)
		c.JSON(http.StatusOK, gin.H{"status": "success", "result": menuResults, "message": "操作成功", "statusCode": 200})
	}
}

func (con SystemUserController) UserTypeList(c *gin.Context) {
	lvdtos := userService.UserTypeList()
	c.JSON(http.StatusOK, gin.H{"result": lvdtos, "message": "操作成功"})
}

func (con SystemUserController) Save(c *gin.Context) {
	tUser := model.TUser{}
	if err := c.ShouldBindJSON(&tUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := userService.SaveUser(tUser)
	c.JSON(http.StatusOK, gin.H{"result": user, "message": "操作成功"})
}

func (con SystemUserController) FllowWorks(c *gin.Context) {
	worksId := c.Query("worksId")

	//userService := service.UserService{UserId: "u_10001"}
	status := userService.FollowWorks(worksId)
	c.JSON(http.StatusOK, gin.H{"result": "success", "message": status, "statusCode": 200})
}

func (con SystemUserController) HisTranInUser(c *gin.Context) {
	result := userService.HisTranInUser()
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}

func (con SystemUserController) FollowStatus(c *gin.Context) {
	worksId := c.Query("worksId")

	//userService := service.UserService{UserId: "u_10001"}
	result := userService.FollowStatus(worksId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}

func (con SystemUserController) GetInfo(c *gin.Context) {
	userId := config.SessionGet("userId").(string)
	result := userService.GetUserById(userId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})

}

func (con SystemUserController) Info(c *gin.Context) {
	userId := config.SessionGet("userId").(string)
	//userId := "20230911160129611666145821"
	result := userService.FindUserById(userId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}

func (con SystemUserController) List(c *gin.Context) {
	result := userService.FindUserList()
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result, "message": "操作成功", "statusCode": 200})
}
