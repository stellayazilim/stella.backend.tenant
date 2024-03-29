package Types

import (
	"gorm.io/gorm"
)

// users shared between tenants
type Session struct {
	gorm.Model
	TenantID string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User     User
	UserID   uint
	Tokens   [2]string `gorm:"type:text[]"`
}
