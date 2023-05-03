package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email    string
	Hash     string
	Posts    []Post    `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
	Messages []Message `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}
