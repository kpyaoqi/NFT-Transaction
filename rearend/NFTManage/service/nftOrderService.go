package service

import (
	"encoding/json"
	"time"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/utils"
	SendMQMessage "yqnft/YQRabbitMQ/SimplePublish"
)

type NftOrderService struct {
}

var nftWorksService NftWorksService
var nftUserChainplatService NftUserChainplatService

func (nftOrderService NftOrderService) SaveGenerateOrderWorksIn(inShelfVO apiResult.InShelfVO) model.NftOrder {
	UserId := config.SessionGet("userId").(string)
	//UserId := nftOrderService.UserId
	worksId := inShelfVO.ID
	state := nftWorksService.WorksState(worksId)
	if state != 0 {
		panic("只能上架,未上架的作品")
	}
	notPayWorksIn := nftOrderService.FindNotPayWorksIn(worksId, UserId)
	if notPayWorksIn.ID != "" {
		return notPayWorksIn
	}
	// 上架商品-未支付
	orderCount := nftOrderService.FindUserOrderCount("", UserId, "", 99, 3, 0)
	if orderCount > 0 {
		panic("请先处理未支付的订单!")
	}
	byId := nftWorksService.FindById(worksId)
	if byId.ID == "" {
		panic("作品不存在!")
	}
	if byId.AuthorID != UserId {
		panic("不是作者")
	}
	//售卖中
	if byId.State == 1 {
		panic("请勿重复操作")
	}
	ok := nftUserChainplatService.HaveAddress(UserId, "xuperChain")
	if !ok {
		panic("用户没有完善百度超级链开放网络地址(address)")
	}
	// TODO 预估gas费,计算上架付款
	money := 0.01
	nftOrder := model.NftOrder{}
	nftOrder.ID = utils.GetTimeNO()
	nftOrder.AssetsID = worksId
	nftOrder.Num = byId.Num
	nftOrder.FromUser = "MakerOne平台"
	nftOrder.CreateTime = time.Now()
	nftOrder.UpdateTime = time.Now()
	nftOrder.ToUser = UserId
	nftOrder.TradeTotal = money
	nftOrder.PayPlat = 0
	nftOrder.TradeType = 3
	nftOrder.Price = 0
	nftOrder.Total = 0
	//TODO 支付二维码
	nftOrder.CodeURL = ""
	nftOrder.Gas = money
	models.DB.Create(&nftOrder)
	return nftOrder
}

func (nftOrderService NftOrderService) FindNotPayWorksIn(worksId string, userId string) model.NftOrder {
	nftOrder := model.NftOrder{}
	models.DB.Table("nft_order").
		Where("toUser=? AND tradeType=? AND assetsId=? AND payState=?", userId, 3, worksId, 0).
		Scan(&nftOrder)
	return nftOrder
}

func (nftOrderService NftOrderService) FindUserOrderCount(fromUser string, toUser string, assetsId string, ePayPlat int, eTradeType int, payState int) int64 {
	nftOrder := model.NftOrder{}
	query := models.DB.Model(&nftOrder)
	query = query.Where("1=1")
	if fromUser != "" {
		query = query.Where("fromUser = ?", fromUser)
	}
	if toUser != "" {
		query = query.Where("toUser = ?", toUser)
	}
	if assetsId != "" {
		query = query.Where("assetsId = ?", assetsId)
	}

	if ePayPlat != 99 {
		query = query.Where("payPlat = ?", ePayPlat)
	}

	if eTradeType != 99 {
		query = query.Where("tradeType = ?", eTradeType)
	}

	if payState != 99 {
		query = query.Where("payState = ?", payState)
	}

	var count int64
	query.Count(&count)
	return count
}

func (nftOrderService NftOrderService) UpdateFindOrderState(orderId string, attach []byte) int64 {
	var state int64
	models.DB.Table("nft_order").
		Select("payState").
		Where("id=?", orderId).
		Scan(&state)
	//已支付
	if state == 1 {
		return 1
	}
	//TODO 上架Gas支付
	// 支付状态
	// 未支付
	// 已支付
	nftOrderService.UpdatePayAfter(orderId, 0, time.Now(), 1, "wx07171446761656c46a0f08c5d97ecd0000", attach)
	return 1
}

func (nftOrderService NftOrderService) UpdatePayAfter(orderId string, ePayPlat int64, txTime time.Time, ePayState int64, transactionId string, attach []byte) bool {
	nftOrder := model.NftOrder{}
	models.DB.Where("id=?", orderId).First(&nftOrder)
	nftOrder.PayPlat = ePayPlat
	nftOrder.PayState = ePayState
	nftOrder.UpdateTime = time.Now()
	nftOrder.TxTime = txTime
	nftOrder.TransactionID = transactionId
	flag := false
	//交易作品
	if nftOrder.TradeType == 5 {
		flag = nftOrderService.UpdatePayTradeWorksAfter(nftOrder)
	} else if nftOrder.TradeType == 3 {
		//上架商品
		flag = nftOrderService.UpdatePayWorksInAfter(nftOrder, attach)
	}
	return flag
}

func (nftOrderService NftOrderService) UpdatePayTradeWorksAfter(order model.NftOrder) bool {
	// 直接修改状态
	go models.DB.Save(&order)
	return true
}

func (nftOrderService NftOrderService) UpdatePayWorksInAfter(order model.NftOrder, attach []byte) bool {
	var inShelfVO apiResult.InShelfVO
	var nftWork model.NftWork
	json.Unmarshal(attach, &inShelfVO)
	// 直接修改状态
	go models.DB.Save(&order)
	nftWork = nftWorksService.FindById(order.AssetsID)
	nftWork.Price = inShelfVO.Price
	nftWork.InTime = inShelfVO.WaitingTime
	nftWork.WaitingTime = inShelfVO.WaitingTime
	nftWork.OutTime = inShelfVO.OutTime
	nftWork.UpdateTime = time.Now()
	nftWork.State = 1
	go models.DB.Save(&nftWork)
	SendMQMessage.ChangeNftWorkByIdMessage(order.AssetsID)
	return true
}
