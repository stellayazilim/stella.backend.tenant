package models

import (
	"gorm.io/gorm"
)

// users shared between tenants
type User struct {
	gorm.Model
	Name        string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Password    string
	Sessions    []Session
	IsValidated bool
	Role        *Role
	RoleID      *uint
}
