package service

import (
	"yqnft/NFTManage/models"
	apiResult "yqnft/NFTManage/models/request_result/api"
)

type NftTypeService struct {
}

func (nftTypeService NftTypeService) GetLabelValuePair() []apiResult.LabelValuePairVO2 {
	labelValuePairVO2s := []apiResult.LabelValuePairVO2{}
	models.DB.Table("nft_type t").
		Select("t.*,t.name AS lable").
		Scan(&labelValuePairVO2s)
	return labelValuePairVO2s
}
