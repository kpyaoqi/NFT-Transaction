package utils

import (
	"strconv"
	"time"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
)

func ToInShelfVO(params map[string]interface{}) apiResult.InShelfVO {
	var inShelfVO apiResult.InShelfVO
	inShelfVO.ID = params["id"].(string)
	inShelfVO.Price, _ = strconv.ParseFloat(params["price"].(string), 64)
	inShelfVO.WaitingTime, _ = time.Parse("2006-01-02", params["waitingTime"].(string))
	inShelfVO.OutTime, _ = time.Parse("2006-01-02", params["outTime"].(string))
	return inShelfVO
}

func ConvertToMenuVO(menus []*model.TMenu) []*apiResult.MenuVO {
	var menuVOs []*apiResult.MenuVO
	for _, menu := range menus {
		menuVO := &apiResult.MenuVO{}
		StructAssign(menuVO, menu)
		tMenus := menu.Children
		if tMenus != nil {
			childMenuVO := ConvertToMenuVO(tMenus)
			menuVO.Children = childMenuVO
		}
		menuVOs = append(menuVOs, menuVO)
	}
	return menuVOs
}
