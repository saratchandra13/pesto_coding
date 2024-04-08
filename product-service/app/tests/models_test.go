package tests

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pesto_coding/product_service/app/models"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func TestConcurrentProductUpdates(t *testing.T) {
	db := setupTestDB()

	productStore := models.NewProductStore(db)

	// Create a new product
	product := &models.Product{
		Name:    "Test Product",
		Price:   10.0,
		Stock:   100,
		Version: 0,
	}
	err := productStore.CreateProduct(product)
	assert.NoError(t, err)

	// Fetch the product once before starting the goroutines
	product = productStore.GetProductById(strconv.Itoa(int(product.ID)))

	// Create a wait group and a channel to signal the goroutines to start
	var wg sync.WaitGroup
	start := make(chan struct{})

	// Simulate concurrent updates
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Wait for the signal to start updating
			<-start

			// Create a new instance of the product for each goroutine
			productCopy := &models.Product{
				Name:    product.Name,
				Price:   20.0,
				Stock:   product.Stock,
				Version: product.Version,
			}

			err := productStore.UpdateProduct(productCopy)

			// We expect an error because the product has been updated by another process
			assert.Error(t, err)
		}()
	}

	// Signal the goroutines to start updating
	close(start)

	// Wait for all goroutines to finish
	wg.Wait()

	// Fetch the product again
	updatedProduct := productStore.GetProductById(strconv.Itoa(int(product.ID)))

	// The product's version should still be 1 because one update would have succeeded
	assert.Equal(t, 1, updatedProduct.Version)
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@/productdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}
	return db
}
