package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID          int
	AccountID   uint
	Title       string
	Description string
	PublishedAt time.Time
	Messages    []Message `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
