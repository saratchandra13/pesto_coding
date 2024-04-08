package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pesto_coding/user_service/app/models"
	"github.com/pesto_coding/user_service/app/services"
	"net/http"
	"strconv"
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
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := uc.UserService.GetUser(uint(userId))
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

	newUser, err := uc.UserService.CreateUser(user.Username, user.Email, user.Password)
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

	isAuthenticated := uc.UserService.AuthenticateUser(user.Username, user.Password)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "logged in"})
}
