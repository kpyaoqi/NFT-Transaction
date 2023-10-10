package service

import "yqnft/NFTManage/models"

type NftUserChainplatService struct {
}

func (nftUserChainplatService NftUserChainplatService) HaveAddress(userId string, code string) bool {
	query := models.DB.Table("nft_user_chainplat")
	query = query.Where("userId = ?", userId)
	query = query.Where("chainPlatName = ?", code)
	query = query.Where("LENGTH(address) > 0")
	var count int64
	query.Count(&count)
	return count > 0
}
