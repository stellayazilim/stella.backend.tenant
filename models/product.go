package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	TenantId    string
	Name        string
	Description string
	Explanation string
	Sku         string
	Specs       map[string]string `gorm:"type:bytea" serializer:"json"`
	Tags        []string          `gorm:"serializer:json"`
	Categories  []*Category       `gorm:"many2many:product_category;"`
}
