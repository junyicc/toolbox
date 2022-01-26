package httpserver

const (
	DefaultPage     = 1
	DefaultPageSize = 20
)

type PageRequest struct {
	Page       int    `form:"page" json:"page"`
	PageSize   int    `form:"page_size" json:"page_size"`
	Sort       string `form:"sort" json:"sort"`
	Descending bool   `form:"descending" json:"descending"`
}

func (p *PageRequest) SetDefault() {
	if p.PageSize == 0 {
		p.PageSize = DefaultPageSize
	}
	if p.Page == 0 {
		p.Page = DefaultPage
	}
}
