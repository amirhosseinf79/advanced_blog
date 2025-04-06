package repositories

import "github.com/amirhosseinf79/advanced_blog/internal/domain/models"

type UserRepository interface {
	GetUserByID(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
	GetAllUsers() ([]models.User, error)
	CheckUserExists(username, email string) (bool, error)
}
