package repositories

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
)

type PostRepository interface {
	CreatePost(post *models.Post) error
	GetPostByID(id string) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id string) error
	GetPostsByAuthorID(authorID int) ([]*models.Post, error)
	GetPostsByFilter(filter dto.PostFilterDTO) ([]*models.Post, error)
}
