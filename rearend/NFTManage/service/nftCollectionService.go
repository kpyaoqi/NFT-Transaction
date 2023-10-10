package service

import (
	"encoding/json"
	"time"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	"yqnft/NFTManage/utils"
	SendMQMessage "yqnft/YQRabbitMQ/SimplePublish"
)

type NftCollectionService struct {
}

var grpc config.GrpcMethod

func (nftCollectionService NftCollectionService) IsExistById(collectionId string) bool {
	nftCollection := nftCollectionService.FindById(collectionId)
	return nftCollection.ID != ""
}

func (nftCollectionService NftCollectionService) FindFairMarketList(res chan<- []apiResult.NFTMarketDTO) {
	nftCollections := []model.NftCollection{}
	results := []apiResult.NFTMarketDTO{}
	userId := "u_10001"
	//userId := config.SessionGet("userId").(string)
	//获取缓存
	results = config.GetNFTMarketDTOListRedis(common.FireCollection + userId)
	if len(results) != 0 {
		res <- results
		return
	}
	models.DB.Order("fire DESC").Limit(4).Find(&nftCollections)
	for _, collection := range nftCollections {
		tUser := model.TUser{}
		//nftWork := model.NftWork{}
		NFTMarketDTO := apiResult.NFTMarketDTO{}
		models.DB.Where("id=?", collection.UserID).Find(&tUser)
		//models.DB.Where("CollectionID=?", collection.ID).Find(&nftWork)
		utils.StructAssign(&NFTMarketDTO, &collection)
		NFTMarketDTO.CollectionID = collection.ID
		NFTMarketDTO.AuthorID = tUser.ID
		NFTMarketDTO.AuthorName = tUser.UserName
		NFTMarketDTO.State = collection.IsIn
		NFTMarketDTO.UserID = userId
		results = append(results, NFTMarketDTO)
	}
	// 存入缓存
	data, _ := json.Marshal(results)
	config.SetRedisWithTime(common.FireCollection+userId, data, 3600)
	res <- results
}

func (nftCollectionService NftCollectionService) NftCollectionDetail(ch chan<- apiResult.CollectionDetailVO, collectionId string) {
	nftCollection := model.NftCollection{}
	CollectionDetailVO := apiResult.CollectionDetailVO{}
	nftWorks := []model.NftWork{}
	tUser := model.TUser{}
	nftCollection = nftCollectionService.FindById(collectionId)
	utils.StructAssign(&CollectionDetailVO, &nftCollection)
	models.DB.Where("collectionId=? AND state!=?", collectionId, 4).Find(&nftWorks)
	models.DB.Where("id = ?", nftCollection.UserID).Find(&tUser)
	CollectionDetailVO.UserName = tUser.UserName
	CollectionDetailVO.Avatar = tUser.Avatar
	CollectionDetailVO.Counts = len(nftWorks)
	ch <- CollectionDetailVO
}

func (nftCollectionService NftCollectionService) FindWorksInCollection(ch chan<- []apiResult.NftDetialWorksVO, collectionId string) {
	nftDetialWorksVOS := []apiResult.NftDetialWorksVO{}
	nftCollection := model.NftCollection{}
	nftWorks := []model.NftWork{}
	nftCollection = nftCollectionService.FindById(collectionId)
	models.DB.Where("collectionId=? AND state!=?", collectionId, 4).Find(&nftWorks)
	for _, nftwork := range nftWorks {
		nftwork_ := nftwork
		var nftDetialWorksVO apiResult.NftDetialWorksVO
		utils.StructAssign(&nftDetialWorksVO, &nftwork_)
		nftDetialWorksVO.CollectionName = nftCollection.Name
		userId := config.SessionGet("userId").(string)
		nftDetialWorksVO.UserID = userId
		nftDetialWorksVOS = append(nftDetialWorksVOS, nftDetialWorksVO)
	}
	ch <- nftDetialWorksVOS
}

func (nftCollectionService NftCollectionService) FindNftMarketList(ch chan<- []apiResult.NFTMarketDTO, collectionName string, typeID string) {
	nftMarketDTOS := []apiResult.NFTMarketDTO{}
	models.DB.Table("nft_collection nc").
		Select("nc.*,u.userName AS authorName,u.id AS authorId").
		Joins("JOIN t_user u ON nc.userId = u.id").
		Where("nc.name LIKE ? AND nc.type=?", "%"+collectionName+"%", typeID).
		Scan(&nftMarketDTOS)
	for i := range nftMarketDTOS {
		userId := config.SessionGet("userId").(string)
		nftMarketDTOS[i].UserID = userId
	}
	ch <- nftMarketDTOS
}

func (nftCollectionService NftCollectionService) UpdateCreateCollection(ch chan<- model.NftCollection, nftCollectionVO apiResult.NftCollectionVO) {
	nftCollection := model.NftCollection{}
	nftCollection = nftCollectionService.FindById(nftCollectionVO.ID)
	if nftCollection.IsIn == 1 {
		panic("合集已上架！不可再修改！")
	}
	userId := config.SessionGet("userId").(string)
	//userId := nftCollectionService.UserId
	if nftCollection.UserID != userId {
		panic("不是作者！")
	}
	if nftCollection.Name != nftCollectionVO.Name {
		if exist := nftCollectionService.IsCollectionNameExist(nftCollectionVO.Name); exist {
			panic("合集名称已存在")
		}
	}
	utils.StructAssign(&nftCollection, &nftCollectionVO)
	nftCollection.UpdateTime = time.Now()
	go models.DB.Updates(&nftCollection)
	SendMQMessage.ChangeNftCollectionByIdMessage(nftCollectionVO.ID)
	ch <- nftCollection
}

func (nftCollectionService NftCollectionService) SaveCreateCollection(ch chan<- model.NftCollection, nftCollectionVO apiResult.NftCollectionVO) {
	var nftCollection model.NftCollection
	userId := config.SessionGet("userId").(string)
	if exist := nftCollectionService.IsCollectionNameExist(nftCollectionVO.Name); exist {
		panic("合集名称已存在")
	}
	utils.StructAssign(&nftCollection, &nftCollectionVO)
	nftCollection.ID = utils.GetTimeNO()
	nftCollection.UserID = userId
	currentTime := time.Now()
	nftCollection.CreateTime = currentTime
	nftCollection.IsIn = 0
	nftCollection.IsCert = 0
	if nftCollection.ChainPlatCert == "" {
		nftCollection.ChainPlatCert = "{}"
	}
	go models.DB.Save(&nftCollection)
	// GRPC调用部署长安链NFT合约
	go grpc.CreateContract(userId, nftCollectionVO)

	nftUserAsset := model.NftUserAsset{}
	nftUserAsset.UserID = userId
	nftUserAsset.AssetsID = nftCollection.ID
	nftUserAsset.ID = utils.GetTimeNO()
	nftUserAsset.SellState = 0
	nftUserAsset.CreateTime = time.Now()
	nftUserAsset.Type = 0
	nftUserAsset.Origin = 0
	go models.DB.Create(&nftUserAsset)
	ch <- nftCollection
}

func (nftCollectionService NftCollectionService) GetLabelValuePair(ch chan<- []apiResult.LabelValuePairVO, userId string) {
	labelValuePairVOs := []apiResult.LabelValuePairVO{}
	if userId != "" {
		models.DB.Table("nft_collection nc").
			Select("nc.*,nc.id AS value,nc.name AS label").
			Where("nc.userId = ?", userId).
			Scan(&labelValuePairVOs)
	} else {
		models.DB.Table("nft_collection nc").
			Select("nc.*,nc.id AS value,nc.name AS label").
			Scan(&labelValuePairVOs)
	}
	ch <- labelValuePairVOs
}

func (nftCollectionService NftCollectionService) IsCollectionNameExist(collectionName string) bool {
	var num int64
	models.DB.Table("nft_collection").Where("name = ?", collectionName).Count(&num)
	if num != 0 {
		return true
	}
	return false
}

func (nftCollectionService NftCollectionService) FindCollectionInUser(ch chan<- []apiResult.NFTMarketDTO) {
	nftCollections := []model.NftCollection{}
	results := []apiResult.NFTMarketDTO{}
	userId := config.SessionGet("userId").(string)
	//userId := nftCollectionService.UserId
	models.DB.Where("userId = ?", userId).Find(&nftCollections)
	for _, collection := range nftCollections {
		tUser := model.TUser{}
		nftWork := model.NftWork{}
		NFTMarketDTO := apiResult.NFTMarketDTO{}
		models.DB.Where("id=?", collection.UserID).Find(&tUser)
		models.DB.Where("CollectionID=?", collection.ID).Find(&nftWork)
		utils.StructAssign(&NFTMarketDTO, &collection)
		NFTMarketDTO.CollectionID = collection.ID
		NFTMarketDTO.AuthorID = tUser.ID
		NFTMarketDTO.AuthorName = tUser.UserName
		NFTMarketDTO.State = collection.IsIn
		NFTMarketDTO.UserID = userId
		results = append(results, NFTMarketDTO)
	}
	ch <- results
}

func (nftCollectionService NftCollectionService) FindById(collectionID string) model.NftCollection {
	//获取缓存
	var nftCollection model.NftCollection
	nftCollection = config.GetNftCollectionRedis(common.CollectionByID + collectionID)
	if nftCollection.ID != "" {
		return nftCollection
	}
	models.DB.Where("id = ?", collectionID).First(&nftCollection)
	data, _ := json.Marshal(nftCollection)
	// 设置缓存
	config.SetRedis(common.CollectionByID+collectionID, data)
	return nftCollection
}
