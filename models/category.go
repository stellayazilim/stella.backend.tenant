package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	TenantID    string
	Name        string
	Description string
	Products    []*Product `gorm:"many2many:product_category;"`
}
