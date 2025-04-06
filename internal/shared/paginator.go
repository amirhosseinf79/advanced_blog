package shared

type Paginator[T any] struct {
	PageSize int64
	Total    int64
	Page     int64
	Data     T
}

type ListResponse[T any] struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	NextPage  int64 `json:"next_page"`
	TotalPage int64 `json:"total_page"`
	Result    T     `json:"result"`
}

func NewPaginator[T any](total int64, page, pageSize int64, data T) *Paginator[T] {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return &Paginator[T]{
		PageSize: pageSize,
		Total:    total,
		Page:     page,
		Data:     data,
	}
}

func (p *Paginator[T]) GetNextPage() int64 {
	if p.Page < p.GetTotalPage() {
		return p.Page + 1
	}
	return p.Page
}

func (p *Paginator[T]) GetTotalPage() int64 {
	if p.Total == 0 {
		return 1
	}
	totalPages := p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		totalPages++
	}
	return totalPages
}

func (p *Paginator[T]) Paginate() ListResponse[T] {
	return ListResponse[T]{
		Total:     p.Total,
		Page:      p.Page,
		NextPage:  p.GetNextPage(),
		TotalPage: p.GetTotalPage(),
		Result:    p.Data,
	}
}
