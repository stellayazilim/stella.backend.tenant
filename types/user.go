package Types

import (
	"gorm.io/gorm"
)

// users shared between tenants
type User struct {
	gorm.Model
	TenantID    string
	Name        string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Password    []byte
	Sessions    []Session
	IsValidated bool
	Role        *Role
	RoleID      *uint
}
