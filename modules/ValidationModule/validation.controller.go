package ValidationModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	"net/http"
)

type IValidationController interface {
	VerifyValidationToken(ctx *gin.Context)
}

type validationController struct {
	validationService IValidationService
}

func ValidationController() IValidationController {
	return &validationController{
		validationService: ValidationService(),
	}
}

func (c validationController) VerifyValidationToken(ctx *gin.Context) {

	id, err := helpers.ConvertUint(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	fmt.Println(id, ctx.Param("token"))
	//isValid := c.validationService.VerifyValidationToken(id, ctx.Param("token"))

	//	fmt.Println(isValid)
}
