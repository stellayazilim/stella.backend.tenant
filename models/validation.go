package models

import (
	"gorm.io/gorm"
)

// users shared between tenants
type Validation struct {
	gorm.Model
	TenantID string
	User     User
	UserID   uint
	Token    string
}
