package service

import (
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/models"
	"yqnft/NFTManage/models/model"
)

type MenuService struct {
}

func (menuService MenuService) FindAllMenuListByQueryBean(data interface{}, page common.Page) []*model.TMenu {
	tMenus := []*model.TMenu{}
	if page.PageNo == 0 {
		page.PageNo = 1
		page.PageSize = 99999
		page.Sort = "DESC"
		page.Order = "sortno"
	}
	//models.DB.Table("t_menu").Order("sortno DESC").Find(&tMenus)
	models.DB.Table("t_menu").Order(page.Order + " " + page.Sort).Find(&tMenus)
	return menuService.MenuList2Tree(tMenus)
}

func (menuService MenuService) MenuList2Tree(menus []*model.TMenu) []*model.TMenu {
	menuMap := make(map[string]*model.TMenu)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}
	var tree []*model.TMenu
	// 遍历菜单列表，构建树形结构
	for _, menu := range menus {
		parent := menuMap[menu.Pid]
		if parent == nil {
			// 没有父节点，说明是根节点
			tree = append(tree, menu)
			continue
		}
		// 有父节点，将当前节点添加到父节点的 Children 列表中
		parent.Children = append(parent.Children, menu)
		menuMap[menu.Pid] = parent
	}
	return tree
}
