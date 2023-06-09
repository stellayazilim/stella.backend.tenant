package CategoryModule

import "github.com/gin-gonic/gin"

func InitCategoryModule(r *gin.RouterGroup) {

	c := CategoryController()
	r.POST("", c.CreateCategory)
	r.GET("", c.GetCategories)
	r.GET(":id", c.GetCategoryById)
	r.PATCH(":id", c.UpdateCategoryById)
	r.DELETE(":id", c.DeleteCategoryById)
}
