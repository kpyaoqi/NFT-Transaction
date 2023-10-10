package common

import (
	"time"
)

// Page 分页对象
type Page struct {
	PageNo          int
	PageSize        int
	TotalCount      int
	PageCount       int
	Sort            string
	Order           string
	FirstPage       bool
	HasPrev         bool
	HasNext         bool
	LastPage        bool
	SelectPageCount bool
	Data            interface{}
	BeginTime       time.Time
	EndTime         time.Time
}

// NewPage 创建一个新的分页对象，默认每页 20 行数据
func NewPage(pageNo int) *Page {
	return &Page{
		PageNo:   pageNo,
		PageSize: 20,
	}
}

// CalculatePageCount 计算总共有多少页
func (p *Page) CalculatePageCount() {
	p.PageCount = (p.TotalCount + p.PageSize - 1) / p.PageSize

	if p.PageNo >= p.PageCount {
		p.LastPage = true
	} else {
		p.HasNext = true
	}

	if p.PageNo > 1 {
		p.HasPrev = true
	} else {
		p.FirstPage = true
	}
}

// IsEmpty 判断是否没有数据
func (p *Page) IsEmpty() bool {
	return p.TotalCount < 1
}
