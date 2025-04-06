package service

import (
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/repositories"
	"github.com/amirhosseinf79/advanced_blog/internal/dto"
	"github.com/amirhosseinf79/advanced_blog/pkg"
)

type userUseCase struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewUserUseCase(
	userRepo repositories.UserRepository,
	tokenRepo repositories.TokenRepository,
) UserService {
	return &userUseCase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (us *userUseCase) RegisterUser(creds dto.UserRegisterRequest) (*models.User, error) {
	var pass string
	var err error
	pass, err = pkg.HashPassword(creds.Password)
	if err != nil {
		return nil, err
	}
	exists, err := us.userRepo.CheckUserExists(creds.Username, creds.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, dto.ErrUserAlreadyExists
	}
	user := models.User{
		Username:  creds.Username,
		Email:     creds.Email,
		Password:  pass,
		FirstName: creds.FirstName,
		LastName:  creds.LastName,
	}
	user, err = us.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *userUseCase) LoginUser(creds dto.UserLoginRequest) (*models.Token, error) {
	user, err := us.userRepo.GetUserByUsername(creds.Username)
	if err != nil {
		return nil, dto.ErrInvalidPassword
	}
	if !user.IsPasswordValid(creds.Password) {
		return nil, dto.ErrInvalidPassword
	}
	token, err := us.tokenRepo.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}
	return token, nil
}
