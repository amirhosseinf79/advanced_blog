package persistence

import (
	"time"

	"github.com/amirhosseinf79/advanced_blog/internal/config"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) repositories.TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) CreateToken(userId uint) (*models.Token, error) {
	var token models.Token
	token.UserId = userId
	if err := r.db.Create(&token).Error; err != nil {
		return &token, err
	}
	return &token, nil
}

func (r *tokenRepository) GetTokenByUUID(uuid string) (*models.Token, error) {
	var token models.Token
	var hoursBefore = time.Now().Add(time.Duration(-config.Config.TokenExpireTime) * time.Hour)
	err := r.db.Preload("User").Where("uuid = ? AND updated_at >= ?", uuid, hoursBefore).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) DeleteToken(uuid string) error {
	if err := r.db.Where("uuid = ?", uuid).Delete(&models.Token{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *tokenRepository) RefreshToken(refreshToken string) (*models.Token, error) {
	var token models.Token
	err := r.db.Where("refresh_token = ?", refreshToken).First(&token).Error
	if err != nil {
		return nil, err
	}
	token.RefreshToken = uuid.New()
	if err := r.db.Save(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}
