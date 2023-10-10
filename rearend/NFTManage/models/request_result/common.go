package result

import "time"

type Page struct {
	PageNo          int         `json:"pageNo"`
	PageSize        int         `json:"pageSize"`
	TotalCount      int         `json:"totalCount"`
	PageCount       int         `json:"pageCount"`
	Sort            string      `json:"sort"`
	Order           string      `json:"order"`
	FirstPage       bool        `json:"firstPage"`
	HasPrev         bool        `json:"hasPrev"`
	HasNext         bool        `json:"hasNext"`
	LastPage        bool        `json:"lastPage"`
	SelectPageCount bool        `json:"selectPageCount"`
	Data            interface{} `json:"data"`
	BeginTime       time.Time   `json:"beginTime"`
	EndTime         time.Time   `json:"endTime"`
}

func NewPage(pageNo, pageSize int) *Page {
	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	return &Page{
		PageNo:   pageNo,
		PageSize: pageSize,
	}
}

func (p *Page) GetFirstResult() int {
	return (p.PageNo - 1) * p.PageSize
}

func (p *Page) IsEmpty() bool {
	return p.TotalCount < 1
}

func (p *Page) IsFirstPage() bool {
	return p.FirstPage
}

func (p *Page) IsHasNext() bool {
	return p.HasNext
}

func (p *Page) IsLastPage() bool {
	return p.LastPage
}
