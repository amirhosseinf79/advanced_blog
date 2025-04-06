package middleware

import (
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"github.com/gofiber/fiber/v3"
)

type paginationMiddleware struct{}

func NewPaginationMiddleware() PaginationMiddleware {
	return &paginationMiddleware{}
}

func (p *paginationMiddleware) PaginationCheck(c fiber.Ctx) error {
	var page dto.PageFilter
	c.Bind().Query(&page)
	if page.Page == 0 {
		page.Page = 1
	}
	if page.PageSize == 0 {
		page.PageSize = 10
	}
	c.Locals("page", page.Page)
	c.Locals("page_size", page.PageSize)
	return c.Next()
}
