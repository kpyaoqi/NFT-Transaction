// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWxMenu = "wx_menu"

// WxMenu mapped from table <wx_menu>
type WxMenu struct {
	ID         string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`                // id
	Name       string    `gorm:"column:name;type:varchar(100)" json:"name"`                      // 菜单名称
	Type       string    `gorm:"column:type;type:varchar(50)" json:"type"`                       // 菜单类型
	Keyword    string    `gorm:"column:keyword;type:varchar(255)" json:"keyword"`                // 消息内容
	URL        string    `gorm:"column:url;type:varchar(255)" json:"url"`                        // 跳转地址
	Pid        string    `gorm:"column:pid;type:varchar(50)" json:"pid"`                         // 上级菜单id
	CreateDate time.Time `gorm:"column:createDate;type:datetime;default:null" json:"createDate"` // 创建时间
	SiteID     string    `gorm:"column:siteId;type:varchar(50)" json:"siteId"`                   // 站点id
	Bak1       string    `gorm:"column:bak1;type:varchar(100)" json:"bak1"`
	Bak2       string    `gorm:"column:bak2;type:varchar(100)" json:"bak2"`
	Bak3       string    `gorm:"column:bak3;type:varchar(100)" json:"bak3"`
	Bak4       string    `gorm:"column:bak4;type:varchar(100)" json:"bak4"`
	Bak5       string    `gorm:"column:bak5;type:varchar(100)" json:"bak5"`
}

// TableName WxMenu's table name
func (*WxMenu) TableName() string {
	return TableNameWxMenu
}
