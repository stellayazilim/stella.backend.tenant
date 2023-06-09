package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	TenantId    string
	Name        string
	Description string
	Perms       []byte `gorm:"type:bytea"`
}
