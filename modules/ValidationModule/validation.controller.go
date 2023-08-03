package ValidationModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/services"
	"net/http"
	"strconv"
)

type IValidationController interface {
	VerifyValidationToken(ctx *gin.Context)
}

type validationController struct {
	validationService Services.IValidationService
}

func ValidationController() IValidationController {
	return &validationController{
		validationService: Services.ValidationService(),
	}
}

func (c validationController) VerifyValidationToken(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

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
