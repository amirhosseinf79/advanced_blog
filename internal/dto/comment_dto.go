package dto

import (
	"time"
)

type CommentCreateDTO struct {
	PostID  uint   `json:"post_id" form:"post_id" validate:"required"`
	Comment string `validate:"required"`
	UserID  uint
}

type CommentFilter struct {
	PostID   uint  `json:"post_id" query:"post_id" validate:"required"`
	Page     int64 `json:"page" query:"page" validate:"gte=1,required"`
	PageSize int64 `json:"page_size" query:"page_size" validate:"gte=1,required"`
}

type CommentResponse struct {
	ID        uint
	Comment   string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
