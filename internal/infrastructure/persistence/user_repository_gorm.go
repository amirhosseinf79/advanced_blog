package persistence

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int) error {
	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CheckUserExists(username, email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Or("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
