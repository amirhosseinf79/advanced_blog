package service

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
)

type CommentService interface {
	AddComment(input dto.CommentCreateDTO) (*models.Comment, error)
	AllPostComments(filter dto.CommentFilter) ([]*models.Comment, error)
}

type PostService interface {
	CreatePost(input dto.PostCreateDTO) (*models.Post, error)
	UpdatePost(input dto.PostUpdateDTO) (*models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	GetAllPosts(filter dto.PostFilterDTO) ([]*models.Post, error)
}

type TokenService interface {
	ValidateToken(token string) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
}

type UserService interface {
	RegisterUser(input dto.UserRegisterRequest) (*models.User, error)
	LoginUser(input dto.UserLoginRequest) (*models.Token, error)
}
