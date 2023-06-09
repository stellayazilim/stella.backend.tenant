package ValidationModule

import (
	"fmt"
	"github.com/lucsky/cuid"
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
)

type IValidationService interface {
	CreateValidationToken(user *models.User)
	VerifyValidationToken(id uint, token string) bool
}

type validationService struct{}

func ValidationService() IValidationService {
	return validationService{}
}

func (v validationService) CreateValidationToken(user *models.User) {
	go func() {
		vl := models.Validation{
			Token: cuid.Slug(),
			User:  *user,
		}
		DatabaseModule.DB.Create(&vl)

		fmt.Println(vl)
	}()
}

func (v validationService) VerifyValidationToken(id uint, token string) bool {
	return true
}
