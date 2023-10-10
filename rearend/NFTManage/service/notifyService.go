package service

import (
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	apiResult "yqnft/NFTManage/models/request_result/api"
)

type NotifyService struct {
}

func (notifyService NotifyService) Message() []apiResult.NotifyVO {
	userId := config.SessionGet("userId").(string)
	//userId := notifyService.UserId
	var notifyVOs []apiResult.NotifyVO
	models.DB.Table("t_notify").
		Where("acceptId=? AND status=?", userId, 0).
		Scan(&notifyVOs)
	return notifyVOs
}
