// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTUserRole = "t_user_role"

// TUserRole mapped from table <t_user_role>
type TUserRole struct {
	ID           string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`       // 编号
	UserID       string    `gorm:"column:userId;type:varchar(50);not null" json:"userId"` // 用户编号
	RoleID       string    `gorm:"column:roleId;type:varchar(50);not null" json:"roleId"` // 角色编号
	CreateTime   time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
	CreateUserID *string   `gorm:"column:createUserId;type:varchar(50)" json:"createUserId"`
	UpdateTime   time.Time `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
	UpdateUserID *string   `gorm:"column:updateUserId;type:varchar(50)" json:"updateUserId"`
	Bak1         *string   `gorm:"column:bak1;type:varchar(100)" json:"bak1"`
	Bak2         *string   `gorm:"column:bak2;type:varchar(100)" json:"bak2"`
	Bak3         *string   `gorm:"column:bak3;type:varchar(100)" json:"bak3"`
	Bak4         *string   `gorm:"column:bak4;type:varchar(100)" json:"bak4"`
	Bak5         *string   `gorm:"column:bak5;type:varchar(100)" json:"bak5"`
}

// TableName TUserRole's table name
func (*TUserRole) TableName() string {
	return TableNameTUserRole
}
