package Types

import (
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/constants/prefixes"
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

// single Role response
type GetRoleResponse struct {
	ID          uint     `json:"id"`
	TenantID    uint     `json:"tenantID"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Perms       string   `json:"perms"`
	Users       []string `json:"users"`
}

// multiple roles response
type GetRolesResponse []GetRoleResponse

func (r *GetRoleResponse) FromRole(data *Role) {
	r.ID = data.ID
	r.TenantID = data.ID
	r.Name = data.Name
	r.Description = data.Name
	r.Perms = fmt.Sprintf("%s/roles/%v/perms", prefixes.BASE_URL, data.ID)

	for _, user := range data.Users {
		r.Users = append(r.Users, fmt.Sprintf("%s/users/%v", prefixes.BASE_URL, user.ID))
	}
}

func (r *GetRolesResponse) FromRoles(data *[]*Role) {
	for _, role := range *data {
		var _role GetRoleResponse
		_role.FromRole(role)
		*r = append(*r, _role)
	}
}

// Get Perms of role response
type RolePermsResponse struct {
	Perms *[]byte `json:"perms"`
}
