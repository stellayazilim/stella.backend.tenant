package Services

import (
	"errors"
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/database"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type IUserService interface {
	Create(user *Types.User) error
	GetUsers(limit int, offset int) ([]Types.User, error)
	GetUserById(id uint) (Types.User, error)
	GetUserByEmail(user *Types.User) error
	UpdateUserById(id uint, user *Types.UserUpdateRequestBody) error
	DeleteUserById(id uint) error
}

type userService struct {
	Database *gorm.DB
}

// constructor
func UserService() IUserService {
	return &userService{
		Database: database.DB.GetDatabase(),
	}
}

func (s userService) Create(user *Types.User) error {
	if err := s.Database.Create(&user).Error; err != nil {
		return errors.New("User already exist")
	}
	return nil
}

func (s userService) GetUsers(limit int, offset int) ([]Types.User, error) {
	var users []Types.User
	result := s.Database.Find(&users).Limit(limit).Offset(offset)
	if result.RowsAffected < 1 || result.Error != nil {
		return users, errors.New("User(s) not found")
	}
	return users, nil
}

func (s userService) GetUserById(id uint) (Types.User, error) {
	var user Types.User
	result := s.Database.Find(&user, id).Limit(0)
	if result.RowsAffected < 1 || result.Error != nil {
		return user, errors.New("User(s) not Found")
	}
	return user, nil
}

func (s userService) GetUserByEmail(user *Types.User) error {
	if err := s.Database.Preload("Role").Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (s userService) UpdateUserById(id uint, user *Types.UserUpdateRequestBody) error {
	data := Types.User{}

	s.Database.Find(&data, id)

	values := reflect.ValueOf(user)

	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}
	return nil
}

func (s userService) DeleteUserById(id uint) error {
	if err := s.Database.Delete(Types.User{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
