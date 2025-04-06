package service

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
)

type postUseCase struct {
	postRepo repositories.PostRepository
}

func NewPostUseCase(postRepo repositories.PostRepository) PostService {
	return &postUseCase{postRepo: postRepo}
}

func (p *postUseCase) CreatePost(post dto.PostCreateDTO) (*models.Post, error) {
	postModel := &models.Post{
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}
	return postModel, p.postRepo.CreatePost(postModel)
}

func (p *postUseCase) UpdatePost(post dto.PostUpdateDTO) (*models.Post, error) {
	postModel := &models.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
	return postModel, p.postRepo.UpdatePost(postModel)
}

func (p *postUseCase) GetPostByID(id string) (*models.Post, error) {
	return p.postRepo.GetPostByID(id)
}

func (p *postUseCase) GetAllPosts(filter dto.PostFilterDTO) ([]*models.Post, error) {
	posts, err := p.postRepo.GetPostsByFilter(filter)
	return posts, err
}
