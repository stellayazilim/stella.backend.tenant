package Types

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

type RoleCreateRequest struct {
	TenantID    string `json:"tenantID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Perms       []byte `json:"perms"`
}

// convert role create object to Role object
func (r *RoleCreateRequest) ConvertToRole() *Role {
	return &Role{
		TenantID:    r.TenantID,
		Name:        r.Name,
		Description: r.Description,
		Perms:       r.Perms,
	}
}
