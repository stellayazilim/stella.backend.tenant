package Services

import (
	"github.com/stellayazilim/stella.backend.tenant/dataase"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
	"log"
)

type IProductService interface {
	CreateProduct(product *Types.Product) error
	GetProducts(limit int, offset int) ([]*Types.Product, error)
	GetProductById(id uint) (Types.Product, error)
	UpdateProductById(id uint, product *Types.Product) error
	DeleteProductById(id uint) error
}

type productService struct {
	Database *gorm.DB
}

func ProductService() IProductService {
	return &productService{
		Database: dataase.DB.GetDatabase(),
	}
}

func (s *productService) CreateProduct(data *Types.Product) error {
	if err := s.Database.Create(data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (s *productService) GetProducts(limit int, offset int) ([]*Types.Product, error) {
	var products []*Types.Product
	s.Database.Preload("Categories").Find(&products).Limit(limit).Offset(offset)
	return products, nil
}

func (s *productService) GetProductById(id uint) (Types.Product, error) {
	var product Types.Product
	if err := s.Database.Preload("Categories").Find(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (s *productService) UpdateProductById(id uint, product *Types.Product) error {
	product.ID = id
	if err := s.Database.Save(product).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s *productService) DeleteProductById(id uint) error {
	if err := s.Database.Delete(Types.Product{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
