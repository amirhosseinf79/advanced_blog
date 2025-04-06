package service

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
)

type commentUseCase struct {
	comentRepo repositories.CommentRepository
}

func NewComentUseCase(comentRepo repositories.CommentRepository) CommentService {
	return &commentUseCase{comentRepo: comentRepo}
}

func (uc *commentUseCase) AddComment(input dto.CommentCreateDTO) (comment *models.Comment, err error) {
	comment = &models.Comment{
		Comment: input.Comment,
		PostID:  input.PostID,
		UserID:  input.UserID,
	}
	err = uc.comentRepo.CreateComment(comment)
	return
}

func (uc *commentUseCase) AllPostComments(filter dto.CommentFilter) (comments []*models.Comment, total int64, err error) {
	comments, total, err = uc.comentRepo.AllPostComments(filter)
	return
}
