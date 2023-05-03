package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	AccountID   uint
	PostID      uint
	Text        string
	PublishedAt time.Time
}
