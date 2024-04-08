package models

import "github.com/jinzhu/gorm"

type ProductStore struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (ps *ProductStore) CreateProduct(product *Product) error {
	return ps.db.Create(product).Error
}

func (ps *ProductStore) GetProductById(id string) *Product {
	var product Product
	ps.db.Find(&product, id)
	return &product
}

func (ps *ProductStore) UpdateProduct(product *Product) error {
	return ps.db.Save(product).Error
}

func (ps *ProductStore) DeleteProduct(id string) error {
	return ps.db.Delete(&Product{}, id).Error
}

func (ps *ProductStore) AttachRole(product *Product, role *Role) {
	ps.db.Model(product).Association("Roles").Append(role)
}
