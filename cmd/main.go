package main

import (
	"log"

	"github.com/amirhosseinf79/advanced_blog/internal/application/api/handlers"
	"github.com/amirhosseinf79/advanced_blog/internal/application/api/middleware"
	"github.com/amirhosseinf79/advanced_blog/internal/config"
	"github.com/amirhosseinf79/advanced_blog/internal/infrastructure/database"
	"github.com/amirhosseinf79/advanced_blog/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/advanced_blog/internal/infrastructure/server"
	"github.com/amirhosseinf79/advanced_blog/internal/service"
)

func main() {
	db, err := database.NewDB(config.Config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	tokenRepo := persistence.NewTokenRepository(db)
	userRepo := persistence.NewUserRepository(db)
	postRepo := persistence.NewPostRepository(db)
	commentRepo := persistence.NewCommentRepository(db)

	tokenUseCase := service.NewTokenUseCase(tokenRepo)
	userUseCase := service.NewUserUseCase(userRepo, tokenRepo)
	postUseCase := service.NewPostUseCase(postRepo)
	commentUseCase := service.NewComentUseCase(commentRepo)

	authMiddleware := middleware.NewTokenMiddleware(tokenUseCase)
	postAuthMiddleware := middleware.NewPostAuthorMiddleware(postUseCase)

	tokenHandler := handlers.NewTokenHandler(tokenUseCase)
	userHandler := handlers.NewUserHandler(userUseCase)
	postHandler := handlers.NewPostHandler(postUseCase)
	commentHandler := handlers.NewCommentHandler(commentUseCase)

	server := server.NewServer(
		authMiddleware,
		postAuthMiddleware,
		tokenHandler,
		userHandler,
		postHandler,
		commentHandler,
	)
	server.SetupRoutes()
	server.Start(":3000")
}
