package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	AccountID    uint
	Account      Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RefreshToken string
}
