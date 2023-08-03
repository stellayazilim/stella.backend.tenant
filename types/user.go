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

type UserCreateRequestBody struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    []byte `json:"password"`
}

func (u *UserCreateRequestBody) ConvertToUser() *User {
	return &User{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
	}
}

// single User responseBody
type UserResponseBody struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsValidated bool   `json:"isValidated"`
	Role        *Role  `json:"role"`
}

// map single User to UserResponseBody
func (u *UserResponseBody) FromUser(data User) {
	u.ID = data.ID
	u.Name = data.Name
	u.Email = data.Email
	u.IsValidated = data.IsValidated
	u.Role = data.Role
}

// multiple users response
type UsersResponseBody []UserResponseBody

// Map []User to UsersResponse
func (e *UsersResponseBody) FromUserSlice(data []User) {
	for _, user := range data {
		*e = append(*e, UserResponseBody{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			IsValidated: user.IsValidated,
			Role:        user.Role,
		})
	}
}

type UserUpdateRequestBody struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Role        *Role  `json:"role"`
}

func (u UserUpdateRequestBody) ConvertToUser() *User {

	return &User{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Role:        u.Role,
	}
}
