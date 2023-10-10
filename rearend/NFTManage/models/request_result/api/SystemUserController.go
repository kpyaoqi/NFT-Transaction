package apiResult

import (
	"time"
)

type MenuVO struct {
	ID         string   `json:"key"`
	Comcodes   []string `json:"comcode" ` // 代码
	Comcode    string
	Title      string    `json:"text" ` // vue使用 meta.title
	Pid        string    `json:"parentKey"`
	Pageurl    string    `json:"pageurl"`
	Code       string    `json:"code" ` // 权限显示key,功能，用于按钮显示判断
	MenuType   int64     `json:"type" ` // 0.功能按钮,1.导航菜单
	Path       string    `json:"path" ` // vue路由地址
	Icon       string    `json:"icon"`
	URL        string    `json:"url" `    // 站外url
	Target     string    `json:"target" ` // 窗口标识
	CreateTime time.Time `json:"createTime"`
	Sortno     int64     `json:"order" `    // 排序,查询时倒叙排列
	Active     int64     `json:"active" `   // 是否有效(0否,1是)
	Children   []*MenuVO `json:"children" ` // 子菜单
}

type LVDTO struct {
	Label string `json:"label" ` // vue使用 meta.title
	Value string `json:"value"`
}
