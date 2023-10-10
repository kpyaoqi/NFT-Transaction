package routes

import (
	"github.com/gin-gonic/gin"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/controller/login"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	loginResult "yqnft/NFTManage/models/request_result/login"
	"yqnft/NFTManage/utils"
)

//	func LoginHandler(c *gin.Context) {
//		// 处理登录逻辑，设置会话数据
//		session := sessions.Default(c)
//		tUser := model.TUser{}
//		var user loginResult.User
//		if err := c.ShouldBindJSON(&user); err != nil {
//			c.JSON(400, gin.H{"error": "Invalid JSON data"})
//			return
//		}
//		passwork := utils.EncoderByMd5With32Bit(user.Password)
//		models.DB.Where("account=?", user.Account).Find(&tUser)
//		if passwork != tUser.Password {
//			c.Err()
//		}
//		c.Set("user", tUser)
//		session.Set("userId", tUser.ID)
//		session.Save()
//		c.Next()
//	}
func LoginHandler(c *gin.Context) {
	// 处理登录逻辑，设置会话数据
	config.InitSession(c)
	tUser := model.TUser{}
	var user loginResult.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	passwork := utils.EncoderByMd5With32Bit(user.Password)
	models.DB.Where("account=?", user.Account).Find(&tUser)
	if passwork != tUser.Password {
		c.Err()
	}
	c.Set("user", tUser)
	config.SessionSet("userId", tUser.ID)
	c.Next()
}

func LoginRoutesInit(router *gin.Engine) {
	loginRouter := router.Group("/nft/api/system")
	loginRouter.POST("/login", LoginHandler, login.LoginController{}.Login)
	loginRouter.POST("/logout", login.LoginController{}.Logout)
}
