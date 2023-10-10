package service

import (
	"time"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
	"yqnft/NFTManage/utils"
)

type NftUserAssetsService struct {
	userId string
}

func (nftUserAssetsService NftUserAssetsService) FindUserAssetsNum(userId string, assetsId string) int64 {
	nftUserAsset := model.NftUserAsset{}
	var count int64
	query := models.DB.Model(&nftUserAsset).Where("userId=?", userId)
	if assetsId != "" {
		query.Where("assetsId = ?", assetsId)
	}
	query.Count(&count)
	return count
}

func (nftUserAssetsService NftUserAssetsService) DeleteAssertByWorkId(userId string, assetsId string) int64 {
	var num int64
	query := models.DB.Where("assetsId=? AND userId=? AND sellState=?", assetsId, userId, 0).Delete(&model.NftUserAsset{})
	num = query.RowsAffected
	return num
}

func (nftUserAssetsService NftUserAssetsService) SaveDefaultNftUserAssetsBatch(nftUserAssets []model.NftUserAsset) bool {
	if len(nftUserAssets) != 0 {
		for i := range nftUserAssets {
			nftUserAssets[i].ID = utils.GetTimeNO()
			nftUserAssets[i].SellState = 0
			nftUserAssets[i].CreateTime = time.Now()
		}
		models.DB.Create(&nftUserAssets)
		return true
	}
	return false
}
