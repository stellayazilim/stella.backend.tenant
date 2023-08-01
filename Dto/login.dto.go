package Dto

import (
	Types "github.com/stellayazilim/stella.backend.tenant/types"
)

type UserLoginDto struct {
	Email    string `json:"email" bind:"required"`
	Password []byte `json:"password" bind:"required"`
}

func (d *UserLoginDto) GetAsUser() Types.User {
	return Types.User{
		Email:    d.Email,
		Password: d.Password,
	}
}
