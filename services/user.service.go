package Services

import (
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strings"
)

type IUserService interface {
	Create(user *types.User) error
	GetUsers(limit int, offset int) ([]types.User, error)
	GetUserById(id uint) (types.User, error)
	GetUserByEmail(user *types.User) error
	UpdateUserById(id uint, user types.User) error
	DeleteUserById(id uint) error
}
type userService struct {
	DB *gorm.DB
}

// constructor
func UserService() IUserService {
	return &userService{
		DB: DatabaseModule.DB,
	}
}

func (s userService) Create(user *types.User) error {
	if err := DatabaseModule.DB.Create(user).Error; err != nil {

		return fmt.Errorf(
			"%v is already exist",
			strings.Replace(regexp.MustCompile("(idx_users_\\w+)").
				FindString(err.Error()), "idx_users_", "", 1))
	}

	return nil
}

func (s userService) GetUsers(limit int, offset int) ([]types.User, error) {
	var users []types.User
	if err := DatabaseModule.DB.Find(&users).Limit(limit).Offset(offset).Error; err == nil {

		return users, nil
	}
	return users, fmt.Errorf("can not retrive user(s)")
}

func (s userService) GetUserById(id uint) (types.User, error) {
	user := types.User{}
	if err := DatabaseModule.DB.Find(&user, id).Limit(0).Error; err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

func (s userService) UpdateUserById(id uint, user types.User) error {
	user.ID = id
	if err := DatabaseModule.DB.Save(&user).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (s userService) GetUserByEmail(user *types.User) error {
	if err := s.DB.Preload("Role").Find(user).Error; err != nil {
		log.Fatal("error happen in user service")
		return err
	}
	return nil
}

func (s userService) DeleteUserById(id uint) error {
	if err := DatabaseModule.DB.Delete(types.User{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
