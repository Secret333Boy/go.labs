package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	AccountID   uint
	Title       string
	Description string
	PublishedAt time.Time
	Messages    []Message `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}
