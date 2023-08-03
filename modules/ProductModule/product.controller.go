package ProductModule

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/modules/ProductModule/DTO"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
	"log"
	"net/http"
	"strconv"
)

type IProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	UpdateProductById(ctx *gin.Context)
	DeleteProductById(ctx *gin.Context)
}
type productController struct {
	productService Services.IProductService
}

func ProductController() IProductController {
	return &productController{
		productService: Services.ProductService(),
	}
}
func (c productController) CreateProduct(ctx *gin.Context) {

}

func (c productController) GetProducts(ctx *gin.Context) {
	serializer := ProductSerializer()
	if products, err := c.productService.GetProducts(-1, -1); err == nil {
		ctx.JSON(http.StatusOK, serializer.SerializeAllFromEntity(products))
	}
}

func (c productController) GetProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	serializer := ProductSerializer()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	product, err := c.productService.GetProductById(uint(id))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, serializer.SerializeFromEntity(&product))
}

func (c productController) UpdateProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Fatal(err, id)
	}

	body := DTO.ProductUpdateDto{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	entity := body.ConvertToEntity()
	if err := c.productService.UpdateProductById(uint(id), &entity); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
	}

	serializer := ProductSerializer()
	ctx.JSON(http.StatusOK, serializer.SerializeFromEntity(&entity))

}

func (c productController) DeleteProductById(ctx *gin.Context) {

}
