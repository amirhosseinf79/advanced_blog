package repositories

import "github.com/amirhosseinf79/advanced_blog/internal/domain/models"

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	AllPostComments(postId uint) ([]*models.Comment, error)
}
