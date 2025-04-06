package service

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
)

type tokenUseCase struct {
	tokenRepo repositories.TokenRepository
}

func NewTokenUseCase(tokenRepo repositories.TokenRepository) TokenService {
	return &tokenUseCase{tokenRepo: tokenRepo}
}

func (uc *tokenUseCase) ValidateToken(token string) (*models.Token, error) {
	tokenModel, err := uc.tokenRepo.GetTokenByUUID(token)
	if err != nil {
		return nil, err
	}
	return tokenModel, nil
}

func (uc *tokenUseCase) RefreshToken(refreshToken string) (*models.Token, error) {
	tokenModel, err := uc.tokenRepo.RefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return tokenModel, nil
}
