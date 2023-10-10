package apiResult

import "time"

// NFTMarketDTO、FindNftMarketListResult
type NFTMarketDTO struct {
	AuthorID     string `gorm:"column:authorId" json:"authorId"`                              // 作者标识=用户标识
	AuthorName   string `gorm:"column:authorName" json:"authorName"`                          // 作者名字
	CollectionID string `gorm:"column:id" json:"collectionId"`                                // 合集标识
	CoverPath    string `gorm:"column:coverPath;type:varchar(100);not null" json:"coverPath"` // 合集封面地址
	Type         int64  `gorm:"column:type;type:int(11);not null" json:"type"`                // 合集类别
	Details      string `gorm:"column:details;type:mediumtext;not null" json:"details"`       // 合集简介
	Name         string `gorm:"column:name;type:varchar(20);not null" json:"name"`            // 合集名称
	LogoPath     string `gorm:"column:logoPath;type:varchar(100);not null" json:"logoPath"`   // logo图地址
	State        int64  `gorm:"column:isIn" json:"state"`                                     // 作品状态(0:未上架,1:(已上架)售卖中,2:已售停,3:已下架,4:已删除)
	UserID       string `gorm:"column:userId" json:"userId"`                                  // 作者标识=用户标识
}

type NftDetialWorksVO struct {
	ID             string  `gorm:"column:id;type:varchar(50);primaryKey" json:"id"` // 唯一标识
	CollectionName string  `json:"collectionName"`
	DataPath       string  `gorm:"column:dataPath;type:varchar(100);not null" json:"dataPath"` // 作品数据地址(图片Base64,上链)
	Name           string  `gorm:"column:name;type:varchar(20);not null" json:"name"`          // 作品名称
	Price          float64 `gorm:"column:price;type:decimal(20,2)" json:"price"`               // 作品价格(初始价格)
	UserID         string  `json:"userId"`
	AuthorID       string  `gorm:"column:authorId;type:varchar(50);not null" json:"authorId"` // 作者标识=用户标识
	State          int64   `gorm:"column:state;type:int(11)" json:"state"`                    // 作品状态(0:未上架,1:(已上架)售卖中,2:已售停,3:已下架,4:已删除)
}

type NftCollectionVO struct {
	ID         string `json:"id"`         // 唯一标识
	LogoPath   string `json:"logoPath"`   // logo图地址
	CoverPath  string `json:"coverPath"`  // 合集封面地址
	BannerPath string `json:"bannerPath"` // banner图地址
	Name       string `json:"name"`       // 合集名称
	Symbol     string `json:"symbol"`     // 合集符号
	Link       string `json:"link"`       // 自定义链接
	Details    string `json:"details"`    // 合集简介
	Type       int64  `json:"type"`       // 合集类别
	Royalties  string `json:"royalties"`  // 合集版税，百分比
	//Royaltie   string     `json:"royaltie"`   // 合集版税，百分比
	InTime  time.Time `json:"inTime"`  // 上架时间
	OutTime time.Time `json:"outTime"` // 下架时间
}

type CollectionDetailVO struct {
	ID            string             `json:"id"`     // 唯一标识
	UserID        string             `json:"userId"` // 用户标识
	UserName      string             `json:"username"`
	Avatar        string             `json:"avatar"`
	Counts        int                `json:"counts"`
	LogoPath      string             `json:"logoPath"`      // logo图地址
	CoverPath     string             `json:"coverPath"`     // 合集封面地址
	BannerPath    string             `json:"bannerPath"`    // banner图地址
	Name          string             `json:"name"`          // 合集名称
	Symbol        string             `json:"symbol"`        // 合集符号
	Link          string             `json:"link"`          // 自定义链接
	Details       string             `json:"details"`       // 合集简介
	Type          int64              `json:"type"`          // 合集类别
	Royalties     string             `json:"royalties"`     // 合集版税，百分比
	ChainPlatCert string             `json:"chainPlatCert"` // 上架后,在不同链平台上链,认证后的数据-json
	IsIn          int64              `json:"isIn"`          // 是否上架 0:否,1:是
	InTime        time.Time          `json:"inTime"`        // 上架时间
	OutTime       time.Time          `json:"outTime"`       // 下架时间
	IsCert        int64              `json:"isCert"`        // 是否平台认证,0:否,1是
	CreateTime    time.Time          `json:"createTime"`
	UpdateTime    time.Time          `json:"updateTime"`
	NftWorks      []NftDetialWorksVO `json:"nftWorksList"`
	SumPrice      float64            `json:"tradesum"`   // 交易总额
	Lowprice      float64            `json:"lowPrice"`   // 地板价
	Buyersnum     int64              `json:"possessnum"` // 拥有者数量
	Bak1          string             `json:"bak1"`
	Bak2          string             `json:"bak2"`
	Bak3          string             `json:"bak3"`
}

type LabelValuePairVO struct {
	Value string `gorm:"column:id" json:"value"`
	Label string `gorm:"column:name" json:"label"`
}
