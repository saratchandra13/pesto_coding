package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pesto_coding/user_service/app/controllers"
	"github.com/pesto_coding/user_service/app/models"
	"github.com/pesto_coding/user_service/app/services"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@/userdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("failed to connect database", err)
		panic(err)
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	userStore := models.NewUserStore(db)
	userService := services.NewUserService(userStore)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	r.GET("/users/:id", userController.GetUser)
	r.POST("/users", userController.CreateUser)
	r.POST("/users/auth", userController.AuthenticateUser)

	err = r.Run()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
