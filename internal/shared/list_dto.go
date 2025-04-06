package shared

type ListResponse[T any] struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	NextPage  int64 `json:"next_page"`
	TotalPage int64 `json:"total_page"`
	Result    T     `json:"result"`
}
