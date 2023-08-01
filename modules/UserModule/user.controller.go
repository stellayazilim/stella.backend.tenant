package UserModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/http"
	"strconv"
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
}

func UserController() IUserController {

	return userController{
		userService:       Services.UserService(),
		validationService: Services.ValidationService(),
		//userSerializer:    UserSerializer(),
	}
}

func (c userController) GetUsers(ctx *gin.Context) {

	if users, err := c.userService.GetUsers(10, 0); err == nil {
		response := Types.UsersResponseBody{}
		response.FromUserSlice(users)
		ctx.JSON(200, response)
	}
}

func (c userController) CreateUser(ctx *gin.Context) {
	body := Types.UserCreateRequestBody{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": helpers.ListOfErrors(err),
		})
		return
	}
	// do not accept weak password
	if err := passwordvalidator.Validate(string(body.Password), 30); err != nil {
		fmt.Println("pasword weak")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := c.userService.Create(body.ConvertToUser()); err != nil {
		// user already exist
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	//response on successfully created
	ctx.Status(http.StatusMultiStatus)
}

func (c userController) GetUserByID(ctx *gin.Context) {
	var id uint
	var response Types.UserResponseBody
	if _id, err := strconv.ParseUint(ctx.Param("id"), 10, 32); err != nil {
		ctx.AbortWithStatus(400)
		return
	} else {
		id = uint(_id)
	}

	if _user, err := c.userService.GetUserById(id); err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	} else {
		response.FromUser(_user)
		ctx.JSON(http.StatusOK, response)
	}

}

func (c userController) UpdateUserByID(ctx *gin.Context) {
	var id uint

	// get id
	if _id, err := strconv.ParseUint(ctx.Param("id"), 10, 32); err != nil {

		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		// abort if id is not uint
		id = uint(_id)
	}

	// get request body
	body := Types.UserUpdateRequestBody{}
	if err := c.userService.UpdateUserById(id, &body); err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	} else {
		ctx.Status(http.StatusOK)
	}

}

func (c userController) DeleteUserByID(ctx *gin.Context) {

}
