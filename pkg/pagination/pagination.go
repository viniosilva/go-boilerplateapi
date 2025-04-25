package pagination

const (
	defaultPage  = 1
	defaultLimit = 10
)

type Params struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p *Params) Normalize() {
	if p.Page <= 0 {
		p.Page = defaultPage
	}
	if p.Limit <= 0 {
		p.Limit = defaultLimit
	}
}

func (p *Params) CalculateOffset() int {
	return (p.Page - 1) * p.Limit
}

type Pagination[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}

func (p *Pagination[T]) SetTotalPages() {
	if p.Limit <= 0 {
		p.TotalPages = 0
		return
	}

	p.TotalPages = int((p.Total / int64(p.Limit)))
	if p.Total%int64(p.Limit) > 0 {
		p.TotalPages++
	}
}

func CopyMetadata[T any, D any](p Pagination[T], data []D) Pagination[D] {
	return Pagination[D]{
		Data:       data,
		Total:      p.Total,
		Page:       p.Page,
		Limit:      p.Limit,
		TotalPages: p.TotalPages,
	}
}
