package apiResult

import "time"

type Attachment struct {
	ID             string    `json:"id"`             // id主键
	OrgID          string    `json:"orgId"`          // 部门id
	OrgTypePathKey string    `json:"orgTypePathKey"` // URL路径中的部门类型,例如 URL路径中的 kjj
	BusinessID     string    `json:"businessId"`     // 业务ID,用于业务关联查询
	AttachmentType int       `json:"attachmentType"` // 附件类型,1.政策附件.2.企业认证文件3.专家认证文件.4.企业个人认证文件.0.其他文件
	FileName       string    `json:"fileName"`       // 附件名称
	FileURL        string    `json:"fileURL"`        // 路径
	URL            string    `json:"url"`            // 路径
	Suffix         string    `json:"suffix"`         // 文件后缀
	FileSize       int64     `json:"fileSize"`       // 文件大小,单位K
	LastDownTime   time.Time `json:"lastDownTime"`   // 最后下载时间
	Sortno         int64     `json:"sortno"`         // 排序,查询时倒叙排列
	Active         int64     `json:"active"`         // 是否有效(0否,1是)
	CreateUser     string    `json:"createUser"`     // 创建者
	CreateTime     time.Time `json:"createTime"`     // 上传时间
	UpdateUser     string    `json:"updateUser"`     // 更新者
	UpdateTime     time.Time `json:"updateTime"`     // 更新时间
	Md5Code        string    `json:"md5Code"`        // 文件的md5值
}
