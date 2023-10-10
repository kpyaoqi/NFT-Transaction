// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLockTable = "lock_table"

// LockTable mapped from table <lock_table>
type LockTable struct {
	RowKey        string `gorm:"column:row_key;type:varchar(128);primaryKey" json:"row_key"`
	Xid           string `gorm:"column:xid;type:varchar(96)" json:"xid"`
	TransactionID int64  `gorm:"column:transaction_id;type:bigint(20)" json:"transaction_id"`
	BranchID      int64  `gorm:"column:branch_id;type:bigint(20);not null" json:"branch_id"`
	ResourceID    string `gorm:"column:resource_id;type:varchar(256)" json:"resource_id"`
	//TableName     string    `gorm:"column:table_name;type:varchar(32)" json:"table_name"`
	Pk          string    `gorm:"column:pk;type:varchar(36)" json:"pk"`
	GmtCreate   time.Time `gorm:"column:gmt_create;type:datetime;default:null" json:"gmt_create"`
	GmtModified time.Time `gorm:"column:gmt_modified;type:datetime;default:null" json:"gmt_modified"`
}

// TableName LockTable's table name
func (*LockTable) TableName() string {
	return TableNameLockTable
}