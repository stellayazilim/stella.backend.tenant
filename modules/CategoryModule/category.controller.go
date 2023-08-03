package CategoryModule

import (
	"github.com/gin-gonic/gin"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
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
	categoryService Services.ICategoryService
}

func CategoryController() ICategoryController {
	return &categoryController{
		categoryService: Services.CategoryService(),
	}
}

func (c categoryController) CreateCategory(ctx *gin.Context) {
	//body := dto.CategoryCreateDto{}
	//if err := ctx.ShouldBindJSON(&body); err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//category := c.categorySerializer.SerializeFromCreateDto(body)
	//if err := c.categoryService.CreateCategory(&category); err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	//	return
	//}

}

func (c categoryController) GetCategories(ctx *gin.Context) {
	////limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 32)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

	//if categories, err := c.categoryService.GetCategories(int(limit)); err == nil {
	//	ctx.JSON(http.StatusOK, c.categorySerializer.SerializeAllFromEntity(categories))
	//}

	ctx.AbortWithStatus(http.StatusNoContent)
}
func (c categoryController) GetCategoryById(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if category, err := c.categoryService.GetCategoryById(uint(id)); err == nil {

		if category.ID == 0 {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		//ctx.JSON(http.StatusOK, c.categorySerializer.SerializeFromEntity(&category))
		return
	}

	ctx.AbortWithStatus(http.StatusInternalServerError)
}
func (c categoryController) UpdateCategoryById(ctx *gin.Context) {

}
func (c categoryController) DeleteCategoryById(ctx *gin.Context) {

}
