package models

import (
	"gorm.io/gorm"
)

// users shared between tenants
type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Password    string
	Sessions    []Session
	IsValidated bool
}
