package database

import (
	"github.com/amirhosseinf79/advanced_blog/internal/config"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config config.Conf) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DbString), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
