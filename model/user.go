package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}
