// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWxPayconfig = "wx_payconfig"

// WxPayconfig mapped from table <wx_payconfig>
type WxPayconfig struct {
	ID              string  `gorm:"column:id;type:varchar(50);primaryKey" json:"id"`
	OrgID           string  `gorm:"column:orgId;type:varchar(50);not null" json:"orgId"`                      // 站点Id
	AppID           string  `gorm:"column:appId;type:varchar(500);not null" json:"appId"`                     // 开发者Id
	Secret          string  `gorm:"column:secret;type:varchar(500);not null" json:"secret"`                   // 应用密钥
	MchID           string  `gorm:"column:mchId;type:varchar(500);not null" json:"mchId"`                     // 微信支付商户号
	Key             string  `gorm:"column:key;type:varchar(500);not null" json:"key"`                         // 交易过程生成签名的密钥，仅保留在商户系统和微信支付后台，不会在网络中传播
	APIV3Key        *string `gorm:"column:apiV3Key;type:varchar(500)" json:"apiV3Key"`                        // V3秘钥
	CertificateFile string  `gorm:"column:certificateFile;type:varchar(500);not null" json:"certificateFile"` // 证书地址,绝对路径
	NotifyURL       *string `gorm:"column:notifyUrl;type:varchar(1000)" json:"notifyUrl"`                     // 通知地址
	SignType        *string `gorm:"column:signType;type:varchar(255)" json:"signType"`                        // 加密方式,MD5和HMAC-SHA256
	Active          int64   `gorm:"column:active;type:int;not null;default:1" json:"active"`                  // 状态 0不可用,1可用
	Bak1            *string `gorm:"column:bak1;type:varchar(100)" json:"bak1"`
	Bak2            *string `gorm:"column:bak2;type:varchar(100)" json:"bak2"`
	Bak3            *string `gorm:"column:bak3;type:varchar(100)" json:"bak3"`
	Bak4            *string `gorm:"column:bak4;type:varchar(100)" json:"bak4"`
	Bak5            *string `gorm:"column:bak5;type:varchar(100)" json:"bak5"`
}

// TableName WxPayconfig's table name
func (*WxPayconfig) TableName() string {
	return TableNameWxPayconfig
}
