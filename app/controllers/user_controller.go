package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lazycoder1995/pesto_coding/app/models"
	"github.com/lazycoder1995/pesto_coding/app/services"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := uc.UserService.GetUser(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserService.CreateUser(user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func (uc *UserController) AuthenticateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAuthenticated := uc.UserService.AuthenticateUser(user.ID, user.Password)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "logged in"})
}
