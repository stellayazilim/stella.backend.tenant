package Services

import (
	"errors"
	"github.com/stellayazilim/stella.backend.tenant/database"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
)

type roleService struct {
	db *gorm.DB
}

type IRoleService interface {
	CreateRole(data *Types.Role) error
	GetRoles() *[]*Types.Role
	GetRoleByID(id uint) (*Types.Role, error)
}

func RoleService() IRoleService {
	return &roleService{
		db: database.DB.GetDatabase(),
	}
}
func (s *roleService) CreateRole(data *Types.Role) error {
	if err := s.db.Model(&Types.Role{}).Create(data).Error; err != nil {
		return errors.New("Role already exist")
	}

	return nil
}
func (s *roleService) GetRoles() *[]*Types.Role {
	roles := &[]*Types.Role{}

	if err := s.db.Preload("Users").Find(&roles).Error; err != nil {
		return roles
	}
	return roles
}

func (s *roleService) GetRoleByID(id uint) (*Types.Role, error) {
	role := Types.Role{}
	if err := s.db.Preload("Users").First(&role, id).Error; err != nil {
		return &role, errors.New("Role does not exist")
	}
	return &role, nil
}
