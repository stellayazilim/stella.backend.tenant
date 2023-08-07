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
	GetPermsOfRoleByID(ctx *gin.Context)
	GetUsersOfRoleByID(ctx *gin.Context)
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
	var roles Types.GetRolesResponse
	roles.FromRoles(c.roleService.GetRoles())
	ctx.JSON(http.StatusOK, roles)
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
		var _role Types.GetRoleResponse
		_role.FromRole(role)
		ctx.JSON(http.StatusOK, _role)
	}
}
func (c *roleController) GetPermsOfRoleByID(ctx *gin.Context) {
	var id uint

	if _id, err := strconv.ParseUint(ctx.Param("id"), 10, 32); err != nil {
		ctx.AbortWithStatus(400)
		return
	} else {
		id = uint(_id)
	}

	c.roleService.GetPermsOfRoleByID(id)

	ctx.JSON(http.StatusOK, Types.RolePermsResponse{
		Perms: c.roleService.GetPermsOfRoleByID(id),
	})
}

func (c *roleController) GetUsersOfRoleByID(ctx *gin.Context) {
	var id uint

	if _id, err := strconv.ParseUint(ctx.Param("id"), 10, 32); err != nil {
		ctx.AbortWithStatus(400)
		return
	} else {
		id = uint(_id)
	}
	users := Types.UsersResponseBody{}
	users.FromUserSlice(c.roleService.GetUsersOfRoleById(id))
	ctx.JSON(http.StatusOK, users)
}
