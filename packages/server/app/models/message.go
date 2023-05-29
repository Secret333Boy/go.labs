package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	ID          int
	AccountID   uint
	PostID      uint
	Text        string
	PublishedAt time.Time
}
