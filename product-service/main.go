package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pesto_coding/product_service/app/controllers"
	"github.com/pesto_coding/product_service/app/models"
	"github.com/pesto_coding/product_service/app/services"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@/productdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Product{})

	productStore := models.NewProductStore(db)
	productService := services.NewProductService(productStore)
	productController := controllers.NewProductController(productService)

	r := gin.Default()

	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.GetProduct)
	r.PUT("/products", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	// add role service and controller
	roleStore := models.NewRoleStore(db)
	roleService := services.NewRoleService(roleStore)
	roleController := controllers.NewRoleController(roleService)

	r.POST("/roles", roleController.CreateRole)
	r.GET("/roles/:id", roleController.GetRole)
	r.POST("/roles/assign", roleController.AssignRoleToProduct)

	r.Run() // listen and serve on 0.0.0.0:8080
}
