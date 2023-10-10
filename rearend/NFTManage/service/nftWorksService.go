package service

import (
	"encoding/json"
	"strings"
	"time"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/utils"
	SendMQMessage "yqnft/YQRabbitMQ/SimplePublish"
)

type NftWorksService struct {
}

var nftUserAssetsService NftUserAssetsService
var nftCollectionService NftCollectionService
var userService UserService

func (nftWorksService NftWorksService) UpdateWorks(nftWorkVO apiResult.NftWorkVO) model.NftWork {
	nftWork := model.NftWork{}
	//userId := nftWorksService.UserId
	userId := config.SessionGet("userId").(string)
	nftWork = nftWorksService.FindById(nftWorkVO.ID)
	if nftWork.State != 0 {
		return model.NftWork{}
	}
	authorIdByWorksId := nftWorksService.FindAuthorIdByWorksId(nftWorkVO.ID)
	if authorIdByWorksId != userId {
		return model.NftWork{}
	}
	utils.StructAssign(&nftWork, &nftWorkVO)
	isExistById := nftCollectionService.IsExistById(nftWorkVO.CollectionID)
	if !isExistById {
		return model.NftWork{}
	}
	nftWorksService.UpdateDefaultValueNftWorks(nftWork)
	// TODO 不出售和售卖中数量不一致
	assetsNum := nftUserAssetsService.FindUserAssetsNum(userId, nftWork.ID)
	deleteNum := nftUserAssetsService.DeleteAssertByWorkId(userId, nftWork.ID)
	if assetsNum != deleteNum {
		return model.NftWork{}
	}
	nftUserAssets := []model.NftUserAsset{}
	for i := 0; i < int(nftWork.Num); i++ {
		nftUserAsset := model.NftUserAsset{}
		nftUserAsset.UserID = userId
		nftUserAsset.AssetsID = nftWork.ID
		nftUserAsset.Type = 1
		nftUserAsset.Origin = 0
		nftUserAssets = append(nftUserAssets, nftUserAsset)
	}
	nftUserAssetsService.SaveDefaultNftUserAssetsBatch(nftUserAssets)
	return nftWork
}

func (nftWorksService NftWorksService) FindAuthorIdByWorksId(id string) string {
	nftWork := nftWorksService.FindById(id)
	return nftWork.AuthorID
}

func (nftWorksService NftWorksService) UpdateDefaultValueNftWorks(nftWork model.NftWork) bool {
	models.DB.Model(&nftWork).UpdateColumn("updateTime", time.Now())
	//清楚缓存
	SendMQMessage.ChangeNftWorkByIdMessage(nftWork.ID)
	return true
}

func (nftWorksService NftWorksService) SaveCreateNftWorks(nftWorkVO apiResult.NftWorkVO) model.NftWork {
	userId := config.SessionGet("userId").(string)
	nftWork := model.NftWork{}
	utils.StructAssign(&nftWork, &nftWorkVO)
	isExistById := nftCollectionService.IsExistById(nftWorkVO.CollectionID)
	if !isExistById {
		return model.NftWork{}
	}
	nftWork.ID = utils.GetTimeNO()
	// 获取用户地址
	user := userService.GetUserById(userId)
	// GRPC调用NFT合约的Mint方法
	collection := nftCollectionService.FindById(nftWorkVO.CollectionID)
	tokenId, _ := grpc.MintNft(nftWork.DataPath, user.Address, collection)
	nftWork.TokenID = tokenId

	nftWorksService.saveDefaultValueNftWorks(nftWork)
	nftUserAssets := []model.NftUserAsset{}
	for i := 0; i < int(nftWork.Num); i++ {
		nftUserAsset := model.NftUserAsset{}
		nftUserAsset.UserID = userId
		nftUserAsset.AssetsID = nftWork.ID
		nftUserAsset.Type = 1
		nftUserAsset.Origin = 0
		nftUserAssets = append(nftUserAssets, nftUserAsset)
	}
	nftUserAssetsService.SaveDefaultNftUserAssetsBatch(nftUserAssets)
	return nftWork
}

func (nftWorksService NftWorksService) saveDefaultValueNftWorks(nftWork model.NftWork) bool {
	userId := config.SessionGet("userId").(string)
	nftWork.CreateTime = time.Now()
	nftWork.AuthorID = userId
	nftWork.State = 0
	models.DB.Save(&nftWork)
	//清楚缓存
	SendMQMessage.ChangeNftWorkByIdMessage(nftWork.ID)
	return true
}

func (nftWorksService NftWorksService) FindWorksById(worksId string) apiResult.NftWorkVO {
	nftWorkVO := apiResult.NftWorkVO{}
	models.DB.Table("nft_works").Where("id = ?", worksId).Scan(&nftWorkVO)
	return nftWorkVO
}

func (nftWorksService NftWorksService) UpdateLogicDel(worksId string) bool {
	nftWork := nftWorksService.FindById(worksId)
	userId := config.SessionGet("userId").(string)
	models.DB.Model(&nftWork).
		Where("authorId = ? AND Id = ? AND state IN (?, ?)", userId, worksId, 0, 1).
		Update("state", 4).
		Update("update_time", time.Now())
	//清楚缓存
	SendMQMessage.ChangeNftWorkByIdMessage(worksId)
	return true
}

func (nftWorksService NftWorksService) NftworksInfo(worksId string) apiResult.WorksInfoVO {
	nftCollection := model.NftCollection{}
	buyWorksInfoResult := apiResult.WorksInfoVO{}
	tUser := model.TUser{}
	nftWork := nftWorksService.FindById(worksId)
	nftCollection = nftCollectionService.FindById(nftWork.CollectionID)
	models.DB.Where("id=?", nftWork.AuthorID).Find(&tUser)
	buyWorksInfoResult.Author = tUser.UserName
	buyWorksInfoResult.Detail = nftWork.Details
	buyWorksInfoResult.CollectionDetail = nftCollection.Details
	buyWorksInfoResult.ChainInfo = nftWork.ChainPlatCert
	return buyWorksInfoResult
}

func (nftWorksService NftWorksService) NftworksBuyDetails(worksId string) apiResult.BuyWorksResult {
	nftWork := nftWorksService.FindById(worksId)
	buyWorksResult := apiResult.BuyWorksResult{}
	nftCollection := nftCollectionService.FindById(nftWork.CollectionID)
	userId := config.SessionGet("userId").(string)
	//根据ID查询用户藏品拥有的数量
	utils.StructAssign(&buyWorksResult, &nftWork)
	buyWorksResult.UserId = userId
	buyWorksResult.CollectionName = nftCollection.Name
	if nftWork.Buyers != "" {
		buyWorksResult.Buyer = strings.Split(nftWork.Buyers, ",")
	}
	var count int64
	models.DB.Table("nft_user_assets").Where("assetsId=? AND userId=?", nftWork.ID, nftWork.AuthorID).Count(&count)
	buyWorksResult.RemainingNum = count
	return buyWorksResult
}

func (nftWorksService NftWorksService) WorksState(worksId string) int64 {
	var state int64
	worksById := nftWorksService.FindById(worksId)
	state = worksById.State
	return state
}

func (nftWorksService NftWorksService) FindById(worksId string) model.NftWork {
	nftWork := model.NftWork{}
	//获取缓存
	nftWork = config.GetNftWorksRedis(common.WorksByID + worksId)
	if nftWork.ID != "" {
		return nftWork
	}
	models.DB.Where("id=?", worksId).Find(&nftWork)
	data, _ := json.Marshal(nftWork)
	// 设置缓存
	config.SetRedis(common.WorksByID+worksId, data)
	return nftWork
}

func (nftWorksService NftWorksService) HisTransForWorks(worksId string) []apiResult.WorksHisTranVO {
	vos := []apiResult.WorksHisTranVO{}
	models.DB.Table("nft_order").
		Where("assetsId=? AND payState = 1", worksId).
		Order("txTime DESC").Limit(5).
		Scan(&vos)
	return vos
}

func (nftWorksService NftWorksService) FindWorksByUserId() []apiResult.NftDetialWorksVO {
	vos := []apiResult.NftDetialWorksVO{}
	userId := config.SessionGet("userId").(string)
	var worksIds []string
	models.DB.Table("nft_user_assets").
		Select("assetsId").
		Where("userId=?", userId).
		Distinct("assetsId").
		Scan(&worksIds)
	for _, workId := range worksIds {
		work := nftWorksService.FindById(workId)
		vo := apiResult.NftDetialWorksVO{}
		utils.StructAssign(&vo, &work)
		collection := nftCollectionService.FindById(work.CollectionID)
		if collection.ID == "" {
			continue
		}
		vo.CollectionName = collection.Name
		vo.UserID = userId
		vos = append(vos, vo)
	}
	return vos
}

func (nftWorksService NftWorksService) FindFollowsInUser() []apiResult.NftDetialWorksVO {
	vos := []apiResult.NftDetialWorksVO{}
	tUserFollows := []model.TUserFollow{}
	nftWorks := []model.NftWork{}
	userId := config.SessionGet("userId").(string)
	models.DB.Table("t_user_follow").
		Where("userID=? AND isValid=1", userId).
		Find(&tUserFollows)
	for _, follow := range tUserFollows {
		work := nftWorksService.FindById(follow.WorksID)
		nftWorks = append(nftWorks, work)
	}
	for _, nftwork := range nftWorks {
		nftwork_ := nftwork
		var nftDetialWorksVO apiResult.NftDetialWorksVO
		utils.StructAssign(&nftDetialWorksVO, &nftwork_)
		collection := nftCollectionService.FindById(nftwork.CollectionID)
		nftDetialWorksVO.CollectionName = collection.Name
		nftDetialWorksVO.UserID = userId
		vos = append(vos, nftDetialWorksVO)
	}
	return vos
}

func (nftWorksService NftWorksService) BuyNftWork(worksId string) string {
	userId := config.SessionGet("userId").(string)
	var works model.NftWork
	works = nftWorksService.FindById(worksId)
	collection := nftCollectionService.FindById(works.CollectionID)
	// GRPC调用长安链NFT合约转移NFT
	fromUser := userService.GetUserById(collection.UserID)
	toUser := userService.GetUserById(userId)
	grpc.TransferNft(works, collection, fromUser.Address, toUser.Address)
	//生成交易订单
	nftOrder := model.NftOrder{}
	nftOrder.ID = utils.GetTimeNO()
	nftOrder.FromUser = collection.UserID
	nftOrder.ToUser = userId
	nftOrder.AssetsID = worksId
	nftOrder.Num = 1
	nftOrder.Price = works.Price
	nftOrder.Total = float64(nftOrder.Num) * nftOrder.Price
	nftOrder.TradeTotal = float64(nftOrder.Num) * nftOrder.Price
	nftOrder.PayState = 1
	nftOrder.TradeType = 5
	nftOrder.CreateTime = time.Now()
	nftOrder.TxTime = time.Now()
	nftOrder.TransactionID = "转账交易ID"
	models.DB.Create(&nftOrder)
	works.Buyers = works.Buyers + "," + userId
	models.DB.Save(&works)
	//清除缓存
	SendMQMessage.ChangeNftWorkByIdMessage(worksId)
	//NFT转移主人
	var nftUserAsset model.NftUserAsset
	models.DB.Where("assetsId=? AND userId=?", worksId, collection.UserID).First(&nftUserAsset)
	nftUserAsset.UserID = userId
	nftUserAsset.UpdateTime = time.Now()
	models.DB.Save(&nftUserAsset)
	return ""
}
