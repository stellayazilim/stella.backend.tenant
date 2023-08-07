package RoleModule

import "github.com/gin-gonic/gin"

func InitRoleModule(r *gin.RouterGroup) {

	c := RoleController()

	r.GET("/", c.GetRoles)
	r.GET("/:id", c.GetRoleByID)
	r.POST("/", c.CreateRole)
}
