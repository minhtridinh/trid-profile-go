package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Bio    string `json:"bio"`
}
