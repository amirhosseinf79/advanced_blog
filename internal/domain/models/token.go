package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	UUID         uuid.UUID `gorm:"primaryKey"`
	RefreshToken uuid.UUID `gorm:"not null;unique"`
	UserId       uint      `gorm:"not null"`
	User         User      `gorm:"foreignKey:UserId"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) error {
	t.UUID = uuid.New()
	t.RefreshToken = uuid.New()
	return nil
}
