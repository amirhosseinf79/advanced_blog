package handlers

import "github.com/gofiber/fiber/v3"

type CommentHandler interface {
	GetComments(c fiber.Ctx) error
	AddComment(c fiber.Ctx) error
}

type PostHandler interface {
	CreatePost(c fiber.Ctx) error
	GetPostByID(c fiber.Ctx) error
	GetAllPosts(c fiber.Ctx) error
	UpdatePost(c fiber.Ctx) error
}

type TokenHandler interface {
	RefreshToken(c fiber.Ctx) error
}

type UserHandler interface {
	RegisterUser(c fiber.Ctx) error
	LoginUser(c fiber.Ctx) error
}
