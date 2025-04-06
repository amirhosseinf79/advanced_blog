package repositories

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	AllPostComments(filter dto.CommentFilter) ([]*models.Comment, int64, error)
}
