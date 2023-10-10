// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTRole = "t_role"

// TRole mapped from table <t_role>
type TRole struct {
	ID           string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`             // 角色ID
	Name         string    `gorm:"column:name;type:varchar(60);not null" json:"name"`           // 角色名称
	Code         string    `gorm:"column:code;type:varchar(255)" json:"code"`                   // 权限编码
	Pid          string    `gorm:"column:pid;type:varchar(50)" json:"pid"`                      // 上级角色ID,暂时不实现
	PrivateOrg   int64     `gorm:"column:privateOrg;type:int(11);not null" json:"privateOrg"`   // 角色的部门是否私有,0否,1是,默认0.当角色私有时,菜单只使用此角色的部门权限,不再扩散到全局角色权限,用于设置特殊的菜单权限.公共权限时部门主管有所管理部门的数据全权限,无论角色是否分配. 私有部门权限时,严格按照配置的数据执行,部门主管可能没有部门权限.
	RoleOrgType  int64     `gorm:"column:roleOrgType;type:int(11);not null" json:"roleOrgType"` // 0自己的数据,1所在部门,2所在部门及子部门数据,3.自定义部门数据,4.全部数据权限
	OrgID        string    `gorm:"column:orgId;type:varchar(50);not null" json:"orgId"`         // 角色的归属部门,只有归属部门的主管和上级主管才可以管理角色,其他人员只能增加归属到角色的人员.不能选择部门或则其他操作,只能添加人员,不然存在提权风险,例如 员工角色下有1000人, 如果给 角色 设置了部门,那这1000人都起效了.
	ShareRole    int64     `gorm:"column:shareRole;type:int(11);not null" json:"shareRole"`     // 角色是否共享,0否 1是,默认0,共享的角色可以被下级部门直接使用,但是下级只能添加人员,不能设置其他属性.共享的角色一般只设置roleOrgType,并不设定部门.
	CreateTime   time.Time `gorm:"column:createTime;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"`
	CreateUserID string    `gorm:"column:createUserId;type:varchar(50)" json:"createUserId"`
	UpdateTime   time.Time `gorm:"column:updateTime;type:datetime;default:CURRENT_TIMESTAMPdefault:null" json:"updateTime"`
	UpdateUserID string    `gorm:"column:updateUserId;type:varchar(50)" json:"updateUserId"`
	Sortno       int64     `gorm:"column:sortno;type:int(11);not null;default:10" json:"sortno"` // 排序,查询时倒叙排列
	Remark       string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                // 备注
	Active       int64     `gorm:"column:active;type:int(11);not null;default:1" json:"active"`  // 是否有效(0否,1是)
	IndexPage    string    `gorm:"column:indexPage;type:varchar(500)" json:"indexPage"`          // 是否有效(0否,1是)
	Bak1         string    `gorm:"column:bak1;type:varchar(100)" json:"bak1"`
	Bak2         string    `gorm:"column:bak2;type:varchar(100)" json:"bak2"`
	Bak3         string    `gorm:"column:bak3;type:varchar(100)" json:"bak3"`
	Bak4         string    `gorm:"column:bak4;type:varchar(100)" json:"bak4"`
	Bak5         string    `gorm:"column:bak5;type:varchar(100)" json:"bak5"`
}

// TableName TRole's table name
func (*TRole) TableName() string {
	return TableNameTRole
}