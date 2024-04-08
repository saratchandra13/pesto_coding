package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"sync"
)

type ProductStore struct {
	db *gorm.DB
	mu sync.Mutex
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
	// Lock the mutex before updating the product
	ps.mu.Lock()
	defer ps.mu.Unlock()

	// Start a new transaction
	tx := ps.db.Begin()

	// Fetch the current product from the database within the transaction
	var currentProduct Product
	tx.Find(&currentProduct, product.ID)

	// Increment the version of the product being updated
	product.Version++

	// Check if the versions match
	if currentProduct.Version != product.Version-1 {
		// If the versions do not match, rollback the transaction and return an error
		tx.Rollback()
		return errors.New("the product has been updated by another process")
	}

	// If the versions match, save the product
	if err := tx.Save(product).Error; err != nil {
		// If the save fails, rollback the transaction and return the error
		tx.Rollback()
		return err
	}

	// If everything is successful, commit the transaction
	return tx.Commit().Error
}

func (ps *ProductStore) DeleteProduct(id string) error {
	return ps.db.Delete(&Product{}, id).Error
}

func (ps *ProductStore) AttachRole(product *Product, role *Role) {
	ps.db.Model(product).Association("Roles").Append(role)
}
