package ProductModule

import (
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"log"
)

type IProductService interface {
	CreateProduct(product models.Product) error
	GetProducts(limit int, offset int) ([]*models.Product, error)
	GetProductById(id uint) (models.Product, error)
	UpdateProductById(id uint, product models.Product) error
	DeleteProductById(id uint) error
}

type productService struct {
}

func ProductService() IProductService {
	return &productService{}
}

func (s productService) CreateProduct(data models.Product) error {
	if err := DatabaseModule.DB.Create(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (s productService) GetProducts(limit int, offset int) ([]*models.Product, error) {
	var products []*models.Product
	DatabaseModule.DB.Preload("Categories").Find(&products).Limit(limit).Offset(offset)
	return products, nil
}

func (s productService) GetProductById(id uint) (models.Product, error) {
	var product models.Product
	if err := DatabaseModule.DB.Preload("Categories").Find(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (s productService) UpdateProductById(id uint, product models.Product) error {
	product.ID = id
	if err := DatabaseModule.DB.Save(&product).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s productService) DeleteProductById(id uint) error {
	if err := DatabaseModule.DB.Delete(models.Product{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
