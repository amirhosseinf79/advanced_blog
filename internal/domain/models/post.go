package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"not null"`
	Content   string         `gorm:"not null"`
	AuthorID  uint           `gorm:"not null"`
	Author    User           `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
