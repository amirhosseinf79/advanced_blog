package dto

type PageFilter struct {
	Page     int64 `json:"page" query:"page" validate:"gte=0"`
	PageSize int64 `json:"page_size" query:"page_size" validate:"gte=0,lte=100"`
}
