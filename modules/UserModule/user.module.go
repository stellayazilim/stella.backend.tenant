package UserModule

import "github.com/gin-gonic/gin"

func InitUserModule(router *gin.RouterGroup) {
	c := UserController()
	// init routers here
	router.POST("", c.CreateUser)
	router.GET("", c.GetUsers)
	router.GET(":id", c.GetUserByID)
	router.PATCH(":id", c.UpdateUserByID)
	router.PATCH(":id/:token", c.UpdateUserByID)
}
