package Types

import (
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

type UserCreateRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    []byte `json:"password"`
}

func (u *UserCreateRequest) ConvertToUser() *User {
	return &User{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
	}
}

// single User response
type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsValidated bool   `json:"isValidated"`
	Role        *Role  `json:"role"`
}

// map single User to UserResponse
func (u *UserResponse) FromUser(data User) {
	u.ID = data.ID
	u.Name = data.Name
	u.Email = data.Email
	u.IsValidated = data.IsValidated
	u.Role = data.Role
}

// multiple users response
type UsersResponse []UserResponse

// Map []User to UsersResponse
func (e *UsersResponse) FromUserSlice(data []User) {
	for _, user := range data {
		*e = append(*e, UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			IsValidated: user.IsValidated,
			Role:        user.Role,
		})
	}
}
