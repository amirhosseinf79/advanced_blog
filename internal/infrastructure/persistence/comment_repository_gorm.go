package persistence

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
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

func (c *commentRepo) AllPostComments(postId uint) (comments []*models.Comment, err error) {
	err = c.db.Omit(clause.Associations).Preload("User").Where("post_id = ?", postId).Find(&comments).Error
	return
}
