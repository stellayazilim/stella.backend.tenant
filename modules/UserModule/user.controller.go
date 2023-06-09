package UserModule

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	"github.com/stellayazilim/stella.backend.tenant/modules/UserModule/dto"
	"github.com/stellayazilim/stella.backend.tenant/modules/ValidationModule"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/http"
)

type IUserController interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUserByID(ctx *gin.Context)
	DeleteUserByID(ctx *gin.Context)
}

type userController struct {
	userService       IUserService
	validationService ValidationModule.IValidationService
	userSerializer    IUserSerializer
}

func UserController() IUserController {

	return userController{
		userService:       UserService(),
		validationService: ValidationModule.ValidationService(),
		userSerializer:    UserSerializer(),
	}
}

func (c userController) GetUsers(ctx *gin.Context) {

	if users, err := c.userService.GetUsers(10, 0); err == nil {

		ctx.JSON(200, c.userSerializer.SerializeAllFromEntity(users))
	}
}

func (c userController) CreateUser(ctx *gin.Context) {
	body := dto.UserCreateDto{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": helpers.ListOfErrors(err),
		})
		return
	}
	// do not accept weak password
	if err := passwordvalidator.Validate(body.Password, 30); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// serialize entity from dto

	user := c.userSerializer.SerializeFromCreateDto(&body)
	if err := c.userService.Create(&user); err != nil {
		// user already exist
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.validationService.CreateValidationToken(&user)
	//response on successfully created
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (c userController) GetUserByID(ctx *gin.Context) {

	id, err := helpers.ConvertUint(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if user, err := c.userService.GetUserById(id); err == nil {
		ctx.JSON(http.StatusOK, c.userSerializer.SerializeFromEntity(user))
	}

}

func (c userController) UpdateUserByID(ctx *gin.Context) {

}

func (c userController) DeleteUserByID(ctx *gin.Context) {

}
