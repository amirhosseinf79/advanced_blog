package database

import (
	"log"
	"os"
	"time"

	"github.com/amirhosseinf79/advanced_blog/internal/config"
	"github.com/amirhosseinf79/advanced_blog/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // Output to stdout
	logger.Config{
		SlowThreshold:             time.Second, // Log queries slower than this threshold
		LogLevel:                  logger.Info, // Log level (Info, Warn, Error, Silent)
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound errors
		Colorful:                  true,        // Enable colorful logs
	},
)

func NewDB(config config.Conf) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DbString), &gorm.Config{Logger: newLogger})
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
