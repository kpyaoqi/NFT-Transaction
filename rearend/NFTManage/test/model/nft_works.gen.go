// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNftWork = "nft_works"

// NftWork mapped from table <nft_works>
type NftWork struct {
	ID            string     `gorm:"column:id;type:varchar(50);primaryKey" json:"id"` // 唯一标识
	TokenID       *int64     `gorm:"column:tokenId;type:int" json:"tokenId"`
	AuthorID      string     `gorm:"column:authorId;type:varchar(50);not null" json:"authorId"`         // 作者标识=用户标识
	CollectionID  string     `gorm:"column:collectionId;type:varchar(50);not null" json:"collectionId"` // 合集标识
	DataPath      string     `gorm:"column:dataPath;type:varchar(100);not null" json:"dataPath"`        // 作品数据地址(图片Base64,上链)
	Name          string     `gorm:"column:name;type:varchar(20);not null" json:"name"`                 // 作品名称
	Type          int64      `gorm:"column:type;type:int;not null" json:"type"`                         // 作品类型
	Link          *string    `gorm:"column:link;type:varchar(300)" json:"link"`                         // 自定义链接
	Details       string     `gorm:"column:details;type:mediumtext;not null" json:"details"`            // 作品简介(富文本)
	Price         *float64   `gorm:"column:price;type:decimal(20,2)" json:"price"`                      // 作品价格(初始价格)
	Num           int64      `gorm:"column:num;type:int;not null" json:"num"`                           // 作者发行数量
	Buyers        *string    `gorm:"column:buyers;type:varchar(5000)" json:"buyers"`                    // 作品购买者(用户标识,标识之间使用逗号隔开)
	InTime        *time.Time `gorm:"column:inTime;type:datetime" json:"inTime"`                         // 上架时间
	OutTime       *time.Time `gorm:"column:outTime;type:datetime" json:"outTime"`                       // 下架时间
	WaitingTime   *time.Time `gorm:"column:waitingTime;type:datetime" json:"waitingTime"`               // 发布等待期,日期之后才可购买
	ChainPlatCert *string    `gorm:"column:chainPlatCert;type:json" json:"chainPlatCert"`               // 上架后,在不同链平台上链,认证后的数据-json
	State         *int64     `gorm:"column:state;type:int" json:"state"`                                // 作品状态(0:未上架,1:(已上架)售卖中,2:已售停,3:已下架,4:已删除)
	CreateTime    *time.Time `gorm:"column:createTime;type:datetime" json:"createTime"`
	UpdateTime    *time.Time `gorm:"column:updateTime;type:datetime" json:"updateTime"`
	Bak1          *string    `gorm:"column:bak1;type:varchar(255)" json:"bak1"`
	Bak2          *string    `gorm:"column:bak2;type:varchar(255)" json:"bak2"`
	Bak3          *string    `gorm:"column:bak3;type:varchar(255)" json:"bak3"`
}

// TableName NftWork's table name
func (*NftWork) TableName() string {
	return TableNameNftWork
}
