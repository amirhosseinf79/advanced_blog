package middleware

import "github.com/gofiber/fiber/v3"

type AuthMiddleware interface {
	TokenAuthMiddleware(c fiber.Ctx) error
}

type PostAuthorMiddleware interface {
	AuthorizePostAuthor(c fiber.Ctx) error
}

type PaginationMiddleware interface {
	PaginationCheck(c fiber.Ctx) error
}
