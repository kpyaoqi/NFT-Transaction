// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTUser = "t_user"

// TUser mapped from table <t_user>
type TUser struct {
	ID           string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`             //
	UserName     *string   `gorm:"column:userName;type:varchar(30)" json:"userName"`            // 姓名
	Account      *string   `gorm:"column:account;type:varchar(50)" json:"account"`              // 账号
	Password     *string   `gorm:"column:password;type:varchar(50)" json:"password"`            // 密码,默认密码123
	Sex          *string   `gorm:"column:sex;type:varchar(5);default:男" json:"sex"`             // 性别
	Mobile       *string   `gorm:"column:mobile;type:varchar(16)" json:"mobile"`                // 手机号码
	Email        *string   `gorm:"column:email;type:varchar(60)" json:"email"`                  // 邮箱
	OpenID       *string   `gorm:"column:openId;type:varchar(200)" json:"openId"`               // 微信openId
	UnionID      *string   `gorm:"column:unionID;type:varchar(200)" json:"unionID"`             // 微信UnionID
	Avatar       *string   `gorm:"column:avatar;type:varchar(2000)" json:"avatar"`              // 头像地址
	UserType     int64     `gorm:"column:userType;type:int;not null;default:1" json:"userType"` // 0会员,1员工,2店长收银,9系统管理员
	CreateTime   time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
	CreateUserID *string   `gorm:"column:createUserId;type:varchar(50)" json:"createUserId"`
	UpdateTime   time.Time `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"`
	UpdateUserID *string   `gorm:"column:updateUserId;type:varchar(50)" json:"updateUserId"`
	Active       int64     `gorm:"column:active;type:int;not null;default:1" json:"active"` // 是否有效(0否,1是)
	Address      string    `gorm:"column:address;type:varchar(42);not null" json:"address"` // 链上地址
	Bak1         *string   `gorm:"column:bak1;type:varchar(100)" json:"bak1"`
	Bak2         *string   `gorm:"column:bak2;type:varchar(100)" json:"bak2"`
	Bak3         *string   `gorm:"column:bak3;type:varchar(100)" json:"bak3"`
	Bak4         *string   `gorm:"column:bak4;type:varchar(100)" json:"bak4"`
	Bak5         *string   `gorm:"column:bak5;type:varchar(100)" json:"bak5"`
}

// TableName TUser's table name
func (*TUser) TableName() string {
	return TableNameTUser
}
