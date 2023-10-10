package apiResult

import "time"

type InShelfVO struct {
	ID          string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"` // 唯一标识
	Price       float64   `gorm:"column:price;type:decimal(20,2)" json:"price"`    // 作品价格(初始价格) 	// 作品价格(初始价格)
	WaitingTime time.Time `gorm:"column:createTime;type:datetime" json:"waitingTime"`
	OutTime     time.Time `gorm:"column:updateTime;type:datetime" json:"outTime"`
}
