package Types

import (
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// users shared between tenants
type User struct {
	gorm.Model
	TenantID    string
	Name        string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Password    []byte
	Sessions    []Session
	IsValidated bool
	Role        *Role
	RoleID      *uint
}

type UserLoginDto struct {
	Email    string `json:"email" bind:"required"`
	Password string `json:"password" bind:"required"`
}

// convert to User model
func (d *UserLoginDto) SerializeToUser() User {
	return User{
		Email: d.Email,
	}
}

// check password of user
func (d *UserLoginDto) ComparePassword() (bool, error) {
	serialized := d.SerializeToUser()
	DatabaseModule.DB.Find(&User{}, &serialized)
	if err := bcrypt.CompareHashAndPassword(serialized.Password, []byte(d.Password)); err != nil {
		return false, err
	}
	return true, nil
}
