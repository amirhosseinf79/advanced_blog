package middleware

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/service"
	"github.com/gofiber/fiber/v3"
)

type postAuthorMiddleware struct {
	postUseCase service.PostService
}

func NewPostAuthorMiddleware(postUseCase service.PostService) PostAuthorMiddleware {
	return &postAuthorMiddleware{postUseCase: postUseCase}
}

func (m *postAuthorMiddleware) AuthorizePostAuthor(c fiber.Ctx) error {
	postID := c.Params("id")
	user := c.Locals("user").(models.User)

	post, err := m.postUseCase.GetPostByID(postID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	if post.AuthorID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You don't have permission to perform this action",
		})
	}

	return c.Next()
}
