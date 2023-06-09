package CategoryModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	"github.com/stellayazilim/stella.backend.tenant/common/serializers"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	"net/http"
	"strconv"
)

type ICategoryController interface {
	CreateCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	GetCategoryById(ctx *gin.Context)
	UpdateCategoryById(ctx *gin.Context)
	DeleteCategoryById(ctx *gin.Context)
}

type categoryController struct {
	categoryServcie    ICategoryService
	categorySerializer serializers.ICategorySerializer
}

func CategoryController() ICategoryController {
	return &categoryController{
		categoryServcie:    categoryServcie{},
		categorySerializer: serializers.CreateCategorySerializer(),
	}
}

func (c categoryController) CreateCategory(ctx *gin.Context) {
	body := dto.CategoryCreateDto{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": helpers.ListOfErrors(err),
		})
		return
	}
	category := c.categorySerializer.SerializeFromCreateDto(body)
	if err := c.categoryServcie.CreateCategory(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"message":  "created",
			"category": fmt.Sprintf("http://localhost:8080/categories/%v", category.ID),
		},
	})
}

func (c categoryController) GetCategories(ctx *gin.Context) {
	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if categories, err := c.categoryServcie.GetCategories(int(limit)); err == nil {
		ctx.JSON(http.StatusOK, c.categorySerializer.SerializeAllFromEntity(categories))
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
func (c categoryController) GetCategoryById(ctx *gin.Context) {
	id, err := helpers.ConvertUint(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if category, err := c.categoryServcie.GetCategoryById(id); err == nil {

		if category.ID == 0 {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusOK, c.categorySerializer.SerializeFromEntity(&category))
		return
	}

	ctx.AbortWithStatus(http.StatusInternalServerError)
}
func (c categoryController) UpdateCategoryById(ctx *gin.Context) {

}
func (c categoryController) DeleteCategoryById(ctx *gin.Context) {

}
