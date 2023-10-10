package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/controller/api"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	loginResult "yqnft/NFTManage/models/request_result/login"
	"yqnft/NFTManage/service"
	"yqnft/NFTManage/utils"
)

type LoginController struct {
	api.SystemUserController
}

var userService service.UserService

func (con LoginController) Login(c *gin.Context) {
	user, _ := c.Get("user")
	tUser, _ := user.(model.TUser)
	loginUser := loginResult.LoginUser{}
	utils.StructAssign(&loginUser, &tUser)
	tRoles := []model.TRole{}
	models.DB.Table("t_role r").
		Select("r.*").
		Joins("JOIN t_user_role re ON re.roleId = r.id").
		Where("re.userId = ?", tUser.ID).
		Where("r.active = ?", 1).
		Order("r.privateOrg DESC, r.sortno DESC").
		Find(&tRoles)
	c.Set("fromLogin", true)
	c.Set("userId", tUser.ID)
	con.Menu(c)
	menuResults, _ := c.Get("menuResults")
	loginResult := loginResult.LoginResult{"", loginUser, false}
	listMenuCodes := []string{"DEFAULT"}
	codeMap := make(map[string]string)
	if resultMap, ok := menuResults.([]apiResult.MenuVO); ok {
		for _, menu := range resultMap {
			var pageURL string
			if menu.Pageurl != "" {
				pageURL = menu.Pageurl
			}
			if menu.Code != "" {
				codeMap[menu.Code] = menu.Pageurl
			}
			listMenuCodes = append(listMenuCodes, pageURL)
		}
	}
	loginResult.User.Codes = listMenuCodes
	loginResult.User.CodeMap = codeMap
	loginResult.User.RoleName = tRoles[0].Name
	loginResult.User.RoleCode = tRoles[0].Code
	loginResult.User.IndexPage = tRoles[0].IndexPage
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": loginResult, "message": nil, "statusCode": 200, "map": nil, "page": nil})
}

func (con LoginController) Logout(c *gin.Context) {
	config.SessionSet("userId", "20230908095458693536608783")
	//utils.SessionDel("userId")
	c.JSON(http.StatusOK, gin.H{"status": "success", "statusCode": 200})
}
