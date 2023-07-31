package UserModule

import (
	"github.com/stellayazilim/stella.backend.tenant/modules/UserModule/dto"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserSerializer interface {
	SerializeAllFromEntity(users []types.User) []userSerializer
	SerializeFromEntity(user types.User) userSerializer
	SerializeFromCreateDto(dto *dto.UserCreateDto) types.User
}
type userSerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func UserSerializer() IUserSerializer {
	return &userSerializer{}
}

func (u userSerializer) SerializeAllFromEntity(users []types.User) []userSerializer {
	var serialized []userSerializer
	for _, user := range users {
		serialized = append(serialized, u.SerializeFromEntity(user))
	}
	return serialized
}
func (u userSerializer) SerializeFromEntity(user types.User) userSerializer {
	return userSerializer{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u userSerializer) SerializeFromCreateDto(dto *dto.UserCreateDto) types.User {
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 16)

	if err != nil {
		log.Fatal(err)
	}
	return types.User{
		Name:        dto.Name,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Password:    password,
	}
}
