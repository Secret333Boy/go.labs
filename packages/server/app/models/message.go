package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	AccountID   uint
	PostID      uint
	Account     Account `gorm:"constraint:OnDelete:CASCADE;"`
	Post        Post    `gorm:"constraint:OnDelete:CASCADE;"`
	Text        string
	PublishedAt time.Time
	//TODO:fix relation
}
