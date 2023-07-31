package Services

import (
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"log"
)

type IProductService interface {
	CreateProduct(product types.Product) error
	GetProducts(limit int, offset int) ([]*types.Product, error)
	GetProductById(id uint) (types.Product, error)
	UpdateProductById(id uint, product *types.Product) error
	DeleteProductById(id uint) error
}

type productService struct {
}

func ProductService() IProductService {
	return &productService{}
}

func (s productService) CreateProduct(data types.Product) error {
	if err := DatabaseModule.DB.Create(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (s productService) GetProducts(limit int, offset int) ([]*types.Product, error) {
	var products []*types.Product
	DatabaseModule.DB.Preload("Categories").Find(&products).Limit(limit).Offset(offset)
	return products, nil
}

func (s productService) GetProductById(id uint) (types.Product, error) {
	var product types.Product
	if err := DatabaseModule.DB.Preload("Categories").Find(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (s productService) UpdateProductById(id uint, product *types.Product) error {
	product.ID = id
	if err := DatabaseModule.DB.Save(product).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s productService) DeleteProductById(id uint) error {
	if err := DatabaseModule.DB.Delete(types.Product{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
