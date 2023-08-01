package UserModule

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
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
	userService       Services.IUserService
	validationService Services.IValidationService
	//userSerializer    IUserSerializer
}

func UserController() IUserController {

	return userController{
		userService:       Services.UserService(),
		validationService: Services.ValidationService(),
		//userSerializer:    UserSerializer(),
	}
}

func (c userController) GetUsers(ctx *gin.Context) {

	//if users, err := c.userService.GetUsers(10, 0); err == nil {

	//ctx.JSON(200, c.userSerializer.SerializeAllFromEntity(users))
	//}
}

func (c userController) CreateUser(ctx *gin.Context) {
	body := Types.UserCreateRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": helpers.ListOfErrors(err),
		})
		return
	}
	// do not accept weak password
	if err := passwordvalidator.Validate(string(body.Password), 30); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := c.userService.Create(body.ConvertToUser()); err != nil {
		// user already exist
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// create validation token
	//c.validationService.CreateValidationToken(&user)

	// send validation token with email

	//response on successfully created
	ctx.Status(http.StatusMultiStatus)
}

func (c userController) GetUserByID(ctx *gin.Context) {

	//id, err := helpers.ConvertUint(ctx.Param("id"))

	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//}

	//if user, err := c.userService.GetUserById(id); err == nil {
	//ctx.JSON(http.StatusOK, c.userSerializer.SerializeFromEntity(user))
	//}

}

func (c userController) UpdateUserByID(ctx *gin.Context) {

}

func (c userController) DeleteUserByID(ctx *gin.Context) {

}
