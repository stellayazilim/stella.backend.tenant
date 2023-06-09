package models

import (
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	TenantId    string
	Maintanence bool
	Bill        float32
	Storage     [2]float32 `gorm:"type:float[]"`
}
