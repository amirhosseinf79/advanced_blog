package persistence

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repositories.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) CreatePost(post *models.Post) error {
	return r.db.Omit(clause.Associations).Create(post).Error
}

func (r *postRepository) GetPostByID(id string) (*models.Post, error) {
	post := &models.Post{}
	err := r.db.Preload("Author").First(post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) GetAllPosts() ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.db.Preload("Author").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) UpdatePost(post *models.Post) error {
	return r.db.Omit(clause.Associations).Save(post).Error
}

func (r *postRepository) DeletePost(id string) error {
	return r.db.Delete(&models.Post{}, "id = ?", id).Error
}

func (r *postRepository) GetPostsByAuthorID(authorID int) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.db.Preload("Author").Where("author_id = ?", authorID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetPostsByFilter(filter dto.PostFilterDTO) (posts []*models.Post, total int64, err error) {
	query := r.db.Model(&models.Post{}).Preload("Author")

	if filter.AuthorName != "" {
		query = query.Joins("JOIN users ON posts.author_id = users.id").
			Where("LOWER(users.first_name) LIKE LOWER(?) OR LOWER(users.last_name) LIKE LOWER(?)", "%"+filter.AuthorName+"%", "%"+filter.AuthorName+"%")
	}
	if filter.Title != "" {
		query = query.Where("LOWER(posts.title) LIKE LOWER(?)", "%"+filter.Title+"%")
	}

	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page, pageSize := filter.Page, filter.PageSize
	offset := (page - 1) * pageSize
	err = query.Offset(int(offset)).Limit(int(pageSize)).Find(&posts).Error
	return
}
