package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yqnft/NFTManage/service"
)

type SystemNotifyController struct {
}

var notifyService service.NotifyService

func (systemNotifyController SystemNotifyController) Message(c *gin.Context) {

	notifyser := service.NotifyService{}
	notifyVOS := notifyser.Message()
	c.JSON(http.StatusOK, gin.H{"status": "success", "result": notifyVOS, "message": "操作成功", "statusCode": 200})
}
