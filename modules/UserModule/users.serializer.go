package UserModule

import (
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/UserModule/dto"
)

type IUserSerializer interface {
	SerializeAllFromEntity(users []models.User) []userSerializer
	SerializeFromEntity(user models.User) userSerializer
	SerializeFromCreateDto(dto *dto.UserCreateDto) models.User
}
type userSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func UserSerializer() IUserSerializer {
	return &userSerializer{}
}

func (u userSerializer) SerializeAllFromEntity(users []models.User) []userSerializer {
	var serialized []userSerializer
	for _, user := range users {
		serialized = append(serialized, u.SerializeFromEntity(user))
	}
	return serialized
}
func (u userSerializer) SerializeFromEntity(user models.User) userSerializer {
	return userSerializer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (u userSerializer) SerializeFromCreateDto(dto *dto.UserCreateDto) models.User {
	return models.User{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Password:    dto.Password,
	}
}
