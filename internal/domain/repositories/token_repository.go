package repositories

import "github.com/amirhosseinf79/advanced_blog/internal/domain/models"

type TokenRepository interface {
	CreateToken(userId uint) (*models.Token, error)
	GetTokenByUUID(uuid string) (*models.Token, error)
	RefreshToken(uuid string) (*models.Token, error)
	DeleteToken(uuid string) error
}
