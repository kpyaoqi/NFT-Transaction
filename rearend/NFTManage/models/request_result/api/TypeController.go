package apiResult

type LabelValuePairVO2 struct {
	Label string `gorm:"column:name" json:"label"`
	Value int64  `gorm:"column:value" json:"value"`
}
