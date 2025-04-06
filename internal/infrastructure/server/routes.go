package server

func (s *server) SetupRoutes() {
	api := s.app.Group("/api/v1")

	auth := api.Group("/auth")
	post := api.Group("/posts")
	comments := api.Group("/comments")

	auth.Post("/refreshToken", s.tokenHandler.RefreshToken)
	auth.Post("/register", s.userHandler.RegisterUser)
	auth.Post("/login", s.userHandler.LoginUser)

	comments.Get("/", s.commentHandler.GetComments)

	post.Get("/", s.postHandler.GetAllPosts)
	post.Get("/:id", s.postHandler.GetPostByID)

	post.Use(s.authMiddleware.TokenAuthMiddleware)
	post.Post("/add", s.postHandler.CreatePost)
	comments.Post("/add", s.commentHandler.AddComment)

	post.Use(s.postAuthorMiddleware.AuthorizePostAuthor)
	post.Put("/update/:id", s.postHandler.UpdatePost)
}
