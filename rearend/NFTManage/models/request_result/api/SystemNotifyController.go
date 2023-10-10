package apiResult

import "time"

type NotifyVO struct {
	ID         string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`            // 主键
	Title      string    `gorm:"column:title;type:varchar(100);not null" json:"title"`       // 消息标题
	Content    string    `gorm:"column:content;type:varchar(5000);not null" json:"content"`  // 消息内容
	NodifyTime time.Time `gorm:"column:nodifyTime;type:datetime;not null" json:"nodifyTime"` // 通知时间
	Status     int64     `gorm:"column:status;type:tinyint(1);not null" json:"status"`       // 消息状态(0未读,1已读)
	Type       int64     `gorm:"column:type;type:tinyint(1);not null" json:"type"`           // 消息类型(1系统运行通知,2推广消息)
	URL        string    `gorm:"column:url;type:varchar(500)" json:"url"`                    // 推广消息跳转的链接
}
