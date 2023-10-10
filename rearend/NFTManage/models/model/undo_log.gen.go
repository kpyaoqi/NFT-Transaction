// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUndoLog = "undo_log"

// UndoLog mapped from table <undo_log>
type UndoLog struct {
	ID           int64     `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"` // increment id
	BranchID     int64     `gorm:"column:branch_id;type:bigint(20);not null" json:"branch_id"`        // branch transaction id
	Xid          string    `gorm:"column:xid;type:varchar(100);not null" json:"xid"`                  // global transaction id
	Context      string    `gorm:"column:context;type:varchar(128);not null" json:"context"`          // undo_log context,such as serialization
	RollbackInfo []byte    `gorm:"column:rollback_info;type:longblob;not null" json:"rollback_info"`  // rollback info
	LogStatus    int64     `gorm:"column:log_status;type:int(11);not null" json:"log_status"`         // 0:normal status,1:defense status
	LogCreated   time.Time `gorm:"column:log_created;type:datetime;not null" json:"log_created"`      // create datetime
	LogModified  time.Time `gorm:"column:log_modified;type:datetime;not null" json:"log_modified"`    // modify datetime
}

// TableName UndoLog's table name
func (*UndoLog) TableName() string {
	return TableNameUndoLog
}