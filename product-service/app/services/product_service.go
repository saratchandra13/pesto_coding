package services

import (
	"github.com/pesto_coding/product_service/app/models"
)

type ProductService struct {
	ProductStore *models.ProductStore
}

func NewProductService(productStore *models.ProductStore) *ProductService {
	return &ProductService{
		ProductStore: productStore,
	}
}

func (ps *ProductService) CreateProduct(product *models.Product) error {
	return ps.ProductStore.CreateProduct(product)
}

func (ps *ProductService) GetProduct(id string) *models.Product {
	return ps.ProductStore.GetProductById(id)
}

func (ps *ProductService) UpdateProduct(product *models.Product) error {
	return ps.ProductStore.UpdateProduct(product)
}

func (ps *ProductService) DeleteProduct(id string) error {
	return ps.ProductStore.DeleteProduct(id)
}

func (ps *ProductService) AttachRole(product *models.Product, role *models.Role) {
	ps.ProductStore.AttachRole(product, role)
}
