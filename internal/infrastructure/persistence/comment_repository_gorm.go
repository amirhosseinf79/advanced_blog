package persistence

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repositories.CommentRepository {
	return &commentRepo{db: db}
}

func (c *commentRepo) CreateComment(comment *models.Comment) error {
	return c.db.Create(comment).Error
}

func (c *commentRepo) AllPostComments(filter dto.CommentFilter) (comments []*models.Comment, total int64, err error) {
	query := c.db.Omit(clause.Associations).Preload("User").Where("post_id = ?", filter.PostID)
	err = query.Count(&total).Offset(int(filter.Page - 1)).Limit(int(filter.PageSize)).Find(&comments).Error
	return
}
