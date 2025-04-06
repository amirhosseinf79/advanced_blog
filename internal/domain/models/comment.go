package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	PostID    uint
	UserID    uint
	Comment   string
	Post      Post           `gorm:"foreignKey:PostID;constraint:OnDelete:SET NULL"`
	User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
