package dto

import (
	"time"
)

type PostCreateDTO struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	AuthorID uint
}

type PostUpdateDTO struct {
	ID       uint   `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	AuthorID uint
}

type PostFilterDTO struct {
	Title      string `json:"title" query:"title"`
	AuthorName string `json:"author_name" query:"author_name"`
	Page       int64  `json:"page" query:"page" validate:"gte=1,required"`
	PageSize   int64  `json:"page_size" query:"page_size" validate:"gte=1,required"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostAuthorResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorName string    `json:"author_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
