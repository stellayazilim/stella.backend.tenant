package ProductModule

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	"github.com/stellayazilim/stella.backend.tenant/common/serializers"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	"net/http"
)

type IProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	UpdateProductById(ctx *gin.Context)
	DeleteProductById(ctx *gin.Context)
}

type productController struct {
	productService    IProductService
	productSerializer serializers.IProductSerializer
}

func ProductController() IProductController {
	return &productController{
		productService:    ProductService(),
		productSerializer: serializers.CreateProductSerializer(),
	}
}

func (c productController) CreateProduct(ctx *gin.Context) {
	body := dto.ProductCreateDto{}
	// validate dto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": helpers.ListOfErrors(err),
		})
		return
	}
	// serialize entity from dto
	product := c.productSerializer.SerializeFromCreateDto(&body)

	if err := c.productService.CreateProduct(*product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created",
	})
}

func (c productController) GetProducts(ctx *gin.Context) {

}

func (c productController) GetProductById(ctx *gin.Context) {

}

func (c productController) UpdateProductById(ctx *gin.Context) {

}

func (c productController) DeleteProductById(ctx *gin.Context) {

}
