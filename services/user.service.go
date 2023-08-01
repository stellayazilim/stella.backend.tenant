package Services

import (
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/Database"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strings"
)

type IUserService interface {
	Create(user *Types.User) error
	GetUsers(limit int, offset int) ([]Types.User, error)
	GetUserById(id uint) (Types.User, error)
	GetUserByEmail(user *Types.User) error
	UpdateUserById(id uint, user Types.User) error
	DeleteUserById(id uint) error
}
type userService struct {
	Database *gorm.DB
}

// constructor
func UserService() IUserService {
	return &userService{
		Database: Database.DB.GetDatabase(),
	}
}

func (s userService) Create(user *Types.User) error {
	if err := s.Database.Create(user).Error; err != nil {

		return fmt.Errorf(
			"%v is already exist",
			strings.Replace(regexp.MustCompile("(idx_users_\\w+)").
				FindString(err.Error()), "idx_users_", "", 1))
	}

	return nil
}

func (s userService) GetUsers(limit int, offset int) ([]Types.User, error) {
	var users []Types.User
	if err := s.Database.Find(&users).Limit(limit).Offset(offset).Error; err == nil {

		return users, nil
	}
	return users, fmt.Errorf("can not retrive user(s)")
}

func (s userService) GetUserById(id uint) (Types.User, error) {
	user := Types.User{}
	if err := s.Database.Find(&user, id).Limit(0).Error; err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

func (s userService) UpdateUserById(id uint, user Types.User) error {
	user.ID = id
	if err := s.Database.Save(&user).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (s userService) GetUserByEmail(user *Types.User) error {
	if err := s.Database.Preload("Role").Find(user).Error; err != nil {
		log.Fatal("error happen in user service")
		return err
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
