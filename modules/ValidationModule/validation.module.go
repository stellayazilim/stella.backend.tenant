package ValidationModule

import "github.com/gin-gonic/gin"

func InitValidationModule(router *gin.RouterGroup) {
	c := ValidationController()

	router.PATCH(":id/:token", c.VerifyValidationToken)
	// init routers here
}
