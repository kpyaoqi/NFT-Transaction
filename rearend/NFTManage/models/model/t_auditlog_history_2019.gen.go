// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTAuditlogHistory2019 = "t_auditlog_history_2019"

// TAuditlogHistory2019 mapped from table <t_auditlog_history_2019>
type TAuditlogHistory2019 struct {
	ID               string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`                      // ID
	OperationType    string    `gorm:"column:operationType;type:varchar(50)" json:"operationType"`           // 操作类型
	OperatorName     string    `gorm:"column:operatorName;type:varchar(50)" json:"operatorName"`             // 操作人姓名
	PreValue         string    `gorm:"column:preValue;type:text" json:"preValue"`                            // 旧值
	CurValue         string    `gorm:"column:curValue;type:text" json:"curValue"`                            // 新值
	OperationTime    time.Time `gorm:"column:operationTime;type:datetime;default:null" json:"operationTime"` // 操作时间
	OperationClass   string    `gorm:"column:operationClass;type:varchar(500)" json:"operationClass"`        // 操作类
	OperationClassID string    `gorm:"column:operationClassID;type:varchar(50)" json:"operationClassID"`     // 记录ID
	Bak1             string    `gorm:"column:bak1;type:varchar(100)" json:"bak1"`
	Bak2             string    `gorm:"column:bak2;type:varchar(100)" json:"bak2"`
	Bak3             string    `gorm:"column:bak3;type:varchar(100)" json:"bak3"`
	Bak4             string    `gorm:"column:bak4;type:varchar(100)" json:"bak4"`
	Bak5             string    `gorm:"column:bak5;type:varchar(100)" json:"bak5"`
}

// TableName TAuditlogHistory2019's table name
func (*TAuditlogHistory2019) TableName() string {
	return TableNameTAuditlogHistory2019
}
