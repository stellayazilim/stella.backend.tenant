package Types

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	TenantID    string
	Name        string
	Description string
	Perms       []byte `gorm:"type:bytea"`
}
