package Types

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	TenantID    string
	Name        string
	Description string
	Perms       []byte `gorm:"type:bytea"`
	Product     *Product
	ProductID   uint
}
