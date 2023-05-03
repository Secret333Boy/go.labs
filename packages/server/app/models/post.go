package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	AccountID   uint
	Account     Account `gorm:"constraint:OnDelete:CASCADE;"`
	Title       string
	Description string
	PublishedAt time.Time
	//Tags        []string
}
