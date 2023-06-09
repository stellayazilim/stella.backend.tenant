package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	TenantId    string
	Name        string
	Description string
	Products    []*Product `gorm:"many2many:product_category;"`
}
