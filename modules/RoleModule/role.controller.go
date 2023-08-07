package RoleModule

import (
	"github.com/gin-gonic/gin"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"net/http"
	"strconv"
)

type roleController struct {
	roleService Services.IRoleService
}

type IRoleController interface {
	CreateRole(ctx *gin.Context)
	AddPermsToRole(ctx *gin.Context)
	RemovePermsFromRole(ctx *gin.Context)
	GetRoles(ctx *gin.Context)
	GetRoleByID(ctx *gin.Context)
}

func RoleController() IRoleController {
	return &roleController{
		roleService: Services.RoleService(),
	}
}

func (c *roleController) CreateRole(ctx *gin.Context) {
	body := Types.RoleCreateRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	if err := c.roleService.CreateRole(body.ConvertToRole()); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "Role already exist",
		})
		return
	}
	ctx.AbortWithStatus(http.StatusOK)

}

func (c *roleController) AddPermsToRole(ctx *gin.Context) {

}

func (c *roleController) RemovePermsFromRole(ctx *gin.Context) {

}
func (c *roleController) GetRoles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.roleService.GetRoles())
}

func (c *roleController) GetRoleByID(ctx *gin.Context) {
	var id uint

	if _id, err := strconv.ParseUint(ctx.Param("id"), 10, 32); err != nil {
		ctx.AbortWithStatus(400)
		return
	} else {
		id = uint(_id)
	}

	if role, err := c.roleService.GetRoleByID(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, role)
	}

}
