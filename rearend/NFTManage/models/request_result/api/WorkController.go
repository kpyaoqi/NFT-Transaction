package apiResult

import "time"

type BuyWorksResult struct {
	UserId         string   `json:"userId"`
	AuthorID       string   `json:"authorId"`
	State          int64    `json:"state"`
	DataPath       string   `json:"logoPath"`
	CollectionName string   `json:"collectionName"`
	Name           string   `json:"worksName"`
	ID             string   `json:"worksNum"`
	Buyer          []string `json:"buyer"`
	Price          float64  `json:"price"`
	RemainingNum   int64    `json:"remainingNum"`
	Num            int64    `json:"totalNumber"`
}

type WorksInfoVO struct {
	Author           string `json:"author"`
	Detail           string `json:"detail"`
	CollectionDetail string `json:"collectionDetail"`
	ChainInfo        string `json:"chainInfo"`
}

type NftWorkVO struct {
	ID           string `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`                   // 唯一标识
	CollectionID string `gorm:"column:collectionId;type:varchar(50);not null" json:"collectionId"` // 合集标识
	DataPath     string `gorm:"column:dataPath;type:varchar(100);not null" json:"dataPath"`        // 作品数据地址(图片Base64,上链)
	Name         string `gorm:"column:name;type:varchar(20);not null" json:"name"`                 // 作品名称
	Type         int64  `gorm:"column:type;type:int(11);not null" json:"type"`                     // 作品类型
	Link         string `gorm:"column:link;type:varchar(300)" json:"link"`                         // 自定义链接
	Details      string `gorm:"column:details;type:mediumtext;not null" json:"details"`            // 作品简介(富文本)
	Num          int64  `gorm:"column:num;type:int(11);not null" json:"num"`                       // 作者发行数量
	Nums         string `json:"nums"`                                                              // 作者发行数量
}

type WorksHisTranVO struct {
	ID         string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`                // 唯一标识
	FromUser   string    `gorm:"column:fromUser;type:varchar(50);not null" json:"fromUser"`      // 卖方
	ToUser     string    `gorm:"column:toUser;type:varchar(50);not null" json:"toUser"`          // 买方
	AssetsID   string    `gorm:"column:assetsId;type:varchar(50);not null" json:"assetsId"`      // 交易品标识(作品)
	Price      float64   `gorm:"column:price;type:decimal(20,2);not null" json:"price"`          // 单价
	TradeType  int64     `gorm:"column:tradeType;type:int(11);not null" json:"tradeType"`        // 交易类型(0:充值,1:上架合集,2:下架合集,3:上架商品,4:下架商品,5:交易作品,6:转移作品)
	TxTime     time.Time `gorm:"column:txTime;type:datetime;default:null" json:"txTime"`         // 交易时间
	CreateTime time.Time `gorm:"column:createTime;type:datetime;default:null" json:"createTime"` // 作者发行数量
}
type UserHisTranVO struct {
	ID         string    `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`                // 唯一标识
	FromUser   string    `gorm:"column:fromUser;type:varchar(50);not null" json:"fromUser"`      // 卖方
	ToUser     string    `gorm:"column:toUser;type:varchar(50);not null" json:"toUser"`          // 买方
	AssetsID   string    `gorm:"column:assetsId;type:varchar(50);not null" json:"assetsId"`      // 交易品标识(作品)
	Price      float64   `gorm:"column:price;type:decimal(20,2);not null" json:"price"`          // 单价
	TradeType  int64     `gorm:"column:tradeType;type:int(11);not null" json:"tradeType"`        // 交易类型(0:充值,1:上架合集,2:下架合集,3:上架商品,4:下架商品,5:交易作品,6:转移作品)
	TxTime     time.Time `gorm:"column:txTime;type:datetime;default:null" json:"txTime"`         // 交易时间
	CreateTime time.Time `gorm:"column:createTime;type:datetime;default:null" json:"createTime"` // 作者发行数量
}
