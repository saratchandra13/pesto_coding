package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lazycoder1995/pesto_coding/app/controllers"
	"github.com/lazycoder1995/pesto_coding/app/models"
	"github.com/lazycoder1995/pesto_coding/app/services"
)

func main() {
	userStore := models.NewUserStore()
	userService := services.NewUserService(userStore)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	r.GET("/users/:id", userController.GetUser)
	r.POST("/users", userController.CreateUser)
	r.POST("/users/auth", userController.AuthenticateUser)

	err := r.Run()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
