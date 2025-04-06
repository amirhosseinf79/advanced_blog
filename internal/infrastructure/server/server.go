package server

import (
	"log"

	"github.com/amirhosseinf79/advanced_blog/internal/application/api/handlers"
	"github.com/amirhosseinf79/advanced_blog/internal/application/api/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type server struct {
	app                  *fiber.App
	authMiddleware       middleware.AuthMiddleware
	postAuthorMiddleware middleware.PostAuthorMiddleware
	paginationMiddleware middleware.PaginationMiddleware
	tokenHandler         handlers.TokenHandler
	userHandler          handlers.UserHandler
	postHandler          handlers.PostHandler
	commentHandler       handlers.CommentHandler
}

func NewServer(
	authMiddleware middleware.AuthMiddleware,
	postAuthorMiddleware middleware.PostAuthorMiddleware,
	paginationMiddleware middleware.PaginationMiddleware,
	tokenHandler handlers.TokenHandler,
	userHandler handlers.UserHandler,
	postHandler handlers.PostHandler,
	commentHandler handlers.CommentHandler,
) *server {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(logger.New())

	return &server{
		app:                  app,
		authMiddleware:       authMiddleware,
		postAuthorMiddleware: postAuthorMiddleware,
		paginationMiddleware: paginationMiddleware,
		tokenHandler:         tokenHandler,
		userHandler:          userHandler,
		postHandler:          postHandler,
		commentHandler:       commentHandler,
	}
}

func (s *server) Start(port string) {
	err := s.app.Listen(port, fiber.ListenConfig{
		EnablePrefork: false,
	})
	if err != nil {
		log.Fatalln("Failed to start server: ", err)
	}
}
