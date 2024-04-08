package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pesto_coding/order_service/app/controllers"
	"github.com/pesto_coding/order_service/app/models"
	"github.com/pesto_coding/order_service/app/services"
)

func main() {
	// Initialize the database connection
	db, err := gorm.Open("mysql", "root:password@/orderdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Order{})

	// Create an instance of OrderStore
	orderStore := models.NewOrderStore(db)

	// Create an instance of OrderService
	orderService := services.NewOrderService(orderStore)

	// Create an instance of OrderController
	orderController := controllers.NewOrderController(orderService)

	// Start the Gin server
	r := gin.Default()

	// Define the routes
	r.POST("/orders", services.RateLimitMiddleware(100), orderController.CreateOrder)
	r.PUT("/orders/:order_id", orderController.UpdateOrderStatus)

	// Run the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
