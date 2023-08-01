package Services

import (
	"fmt"
	"github.com/lucsky/cuid"
	"github.com/stellayazilim/stella.backend.tenant/Database"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
)

type IValidationService interface {
	CreateValidationToken(user *Types.User)
	VerifyValidationToken(id uint, token string) bool
}

type validationService struct {
	Database *gorm.DB
}

func ValidationService() IValidationService {
	return &validationService{
		Database: Database.DB.GetDatabase(),
	}
}

func (v *validationService) CreateValidationToken(user *Types.User) {
	go func() {
		vl := Types.Validation{
			Token: cuid.Slug(),
			User:  *user,
		}
		v.Database.Create(&vl)

		fmt.Println(vl)
	}()
}

func (v *validationService) VerifyValidationToken(id uint, token string) bool {
	return true
}
