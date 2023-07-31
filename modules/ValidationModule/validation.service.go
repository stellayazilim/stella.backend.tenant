package ValidationModule

import (
	"fmt"
	"github.com/lucsky/cuid"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"github.com/stellayazilim/stella.backend.tenant/types"
)

type IValidationService interface {
	CreateValidationToken(user *types.User)
	VerifyValidationToken(id uint, token string) bool
}

type validationService struct{}

func ValidationService() IValidationService {
	return validationService{}
}

func (v validationService) CreateValidationToken(user *types.User) {
	go func() {
		vl := types.Validation{
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
