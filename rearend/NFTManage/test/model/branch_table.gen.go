// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBranchTable = "branch_table"

// BranchTable mapped from table <branch_table>
type BranchTable struct {
	BranchID        int64      `gorm:"column:branch_id;type:bigint;primaryKey" json:"branch_id"`
	Xid             string     `gorm:"column:xid;type:varchar(128);not null" json:"xid"`
	TransactionID   *int64     `gorm:"column:transaction_id;type:bigint" json:"transaction_id"`
	ResourceGroupID *string    `gorm:"column:resource_group_id;type:varchar(32)" json:"resource_group_id"`
	ResourceID      *string    `gorm:"column:resource_id;type:varchar(256)" json:"resource_id"`
	BranchType      *string    `gorm:"column:branch_type;type:varchar(8)" json:"branch_type"`
	Status          *int64     `gorm:"column:status;type:tinyint" json:"status"`
	ClientID        *string    `gorm:"column:client_id;type:varchar(64)" json:"client_id"`
	ApplicationData *string    `gorm:"column:application_data;type:varchar(2000)" json:"application_data"`
	GmtCreate       *time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`
	GmtModified     *time.Time `gorm:"column:gmt_modified;type:datetime" json:"gmt_modified"`
}

// TableName BranchTable's table name
func (*BranchTable) TableName() string {
	return TableNameBranchTable
}
