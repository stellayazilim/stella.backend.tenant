package ProductModule

import "github.com/gin-gonic/gin"

func InitProductModule(r *gin.RouterGroup) {

	c := ProductController()

	r.POST("", c.CreateProduct)
	r.GET("", c.GetProducts)
	r.GET(":id", c.GetProductById)
	r.PATCH(":id", c.UpdateProductById)
	r.DELETE(":id", c.DeleteProductById)
}
