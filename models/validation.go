package models

import (
	"gorm.io/gorm"
)

// users shared between tenants
type Validation struct {
	gorm.Model
	User   User
	UserID uint
	Token  string
}
