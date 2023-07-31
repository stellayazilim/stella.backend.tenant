package ProductModule

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	"github.com/stellayazilim/stella.backend.tenant/modules/ProductModule/DTO"
	"log"
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
	productService IProductService
}

func ProductController() IProductController {
	return &productController{
		productService: ProductService(),
	}
}
func (c productController) CreateProduct(ctx *gin.Context) {

	body := DTO.ProductCreateDto{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": helpers.ListOfErrors(err),
		})
		log.Fatal(err)
	}

	if err := c.productService.CreateProduct(body.ConvertToEntity()); err != nil {
		// todo handle error
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		log.Fatal(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created",
	})
}

func (c productController) GetProducts(ctx *gin.Context) {
	serializer := ProductSerializer()
	if products, err := c.productService.GetProducts(-1, -1); err == nil {
		ctx.JSON(http.StatusOK, serializer.SerializeAllFromEntity(products))
	}
}

func (c productController) GetProductById(ctx *gin.Context) {
	id, err := helpers.ConvertUint(ctx.Param("id"))
	serializer := ProductSerializer()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": helpers.ListOfErrors(err),
		})
		log.Fatal(err)
	}
	product, err := c.productService.GetProductById(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, serializer.SerializeFromEntity(&product))
}

func (c productController) UpdateProductById(ctx *gin.Context) {
	id, err := helpers.ConvertUint(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": helpers.ListOfErrors(err),
		})
		log.Fatal(err, id)
	}

	body := DTO.ProductUpdateDto{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": helpers.ListOfErrors(err),
		})
		log.Fatal(err)
	}

	entity := body.ConvertToEntity()
	if err := c.productService.UpdateProductById(id, &entity); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": helpers.ListOfErrors(err),
		})
	}

	serializer := ProductSerializer()
	ctx.JSON(http.StatusOK, serializer.SerializeFromEntity(&entity))

}

func (c productController) DeleteProductById(ctx *gin.Context) {

}
