package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	TenantID    string
	Name        string
	Description string
	Perms       []byte `gorm:"type:bytea"`
	Users       []*User
}
