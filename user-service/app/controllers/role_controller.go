package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pesto_coding/user_service/app/models"
	"github.com/pesto_coding/user_service/app/services"
	"net/http"
)

type RoleController struct {
	RoleService *services.RoleService
	UserService *services.UserService
}

func NewRoleController(roleService *services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRole(c *gin.Context) {
	name := c.Param("name")
	role, err := rc.RoleService.GetRoleByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRole, err := rc.RoleService.CreateRole(role.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": newRole})
}

func (rc *RoleController) AssignRoleToUser(c *gin.Context) {
	var userRole models.UserRole
	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := rc.UserService.GetUser(userRole.UserId)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "user not found"})
		return
	}

	role, err := rc.RoleService.GetRoleById(int(userRole.RoleId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "role not found"})
		return
	}

	err = rc.RoleService.AssignRoleToUser(user, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "role assigned to user"})
}
