// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTPermissionsLog = "t_permissions_log"

// TPermissionsLog mapped from table <t_permissions_log>
type TPermissionsLog struct {
	ID                 string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`                       // 主键
	SiteID             string    `gorm:"column:siteId;type:varchar(50)" json:"siteId"`                          // 站点ID
	ActionType         int64     `gorm:"column:actionType;type:int(11)" json:"actionType"`                      // 操作类型 创建、编辑、删除、启用、禁用
	OperatorUserID     string    `gorm:"column:operatorUserId;type:varchar(50)" json:"operatorUserId"`          // 操作人ID
	OperatorUserName   string    `gorm:"column:operatorUserName;type:varchar(200)" json:"operatorUserName"`     // 操作人当时名称
	OperatorUserRoles  string    `gorm:"column:operatorUserRoles;type:text" json:"operatorUserRoles"`           // 操作人当时所属角色名称，逗号分割
	OperatorObjectType int64     `gorm:"column:operatorObjectType;type:int(11)" json:"operatorObjectType"`      // 操作对象类型
	OperatorObjectID   string    `gorm:"column:operatorObjectId;type:varchar(50)" json:"operatorObjectId"`      // 操作对象ID
	OperatorObjectName string    `gorm:"column:operatorObjectName;type:varchar(200)" json:"operatorObjectName"` // 操作对象当时的名称
	ActionContent      string    `gorm:"column:actionContent;type:longtext" json:"actionContent"`               // 操作内容详情
	CreateUserID       string    `gorm:"column:createUserId;type:varchar(50)" json:"createUserId"`              // 记录创建人
	CreateTime         time.Time `gorm:"column:createTime;type:datetime;default:null" json:"createTime"`        // 记录创建时间
	Bak1               string    `gorm:"column:bak1;type:varchar(200)" json:"bak1"`                             // 备用字段
	Bak2               string    `gorm:"column:bak2;type:varchar(200)" json:"bak2"`                             // 备用字段
	Bak3               string    `gorm:"column:bak3;type:varchar(200)" json:"bak3"`                             // 备用字段
	Bak4               string    `gorm:"column:bak4;type:varchar(200)" json:"bak4"`                             // 备用字段
	Bak5               string    `gorm:"column:bak5;type:varchar(200)" json:"bak5"`                             // 备用字段
}

// TableName TPermissionsLog's table name
func (*TPermissionsLog) TableName() string {
	return TableNameTPermissionsLog
}
