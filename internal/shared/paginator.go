package shared

type paginatore[T any] struct {
	pageSize int64
	total    int64
	page     int64
	data     T
}

type Paginatore[T any] interface {
	Paginate() ListResponse[T]
	getNextPage() int64
	getTotalPage() int64
	Validate() (int64, int64)
}

func NewPaginator[T any](total int64, page, pageSize int64, data T) Paginatore[T] {
	return &paginatore[T]{
		pageSize: pageSize,
		total:    total,
		page:     page,
		data:     data,
	}
}

func (p *paginatore[T]) Validate() (int64, int64) {
	if p.page < 1 {
		p.page = 1
	}
	if p.pageSize < 1 {
		p.pageSize = 10
	}
	return p.page, p.pageSize
}

func (p *paginatore[T]) getNextPage() int64 {
	if p.page < p.getTotalPage() {
		return p.page + 1
	}

	return p.page
}

func (p *paginatore[T]) getTotalPage() int64 {
	t := p.total / p.pageSize
	if t == 0 {
		return 1
	}
	return t
}

func (p *paginatore[T]) Paginate() ListResponse[T] {
	p.Validate()
	return ListResponse[T]{
		Total:     p.total,
		Page:      p.page,
		NextPage:  p.getNextPage(),
		TotalPage: p.getTotalPage(),
		Result:    p.data,
	}
}
