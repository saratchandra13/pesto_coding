package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pesto_coding/product_service/app/models"
	"github.com/pesto_coding/product_service/app/services"
	"net/http"
	"strconv"
)

type RoleController struct {
	RoleService *services.RoleService
}

func NewRoleController(roleService *services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role *models.Role
	var err error
	role, err = rc.RoleService.GetRoleById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	var role *models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role, err := rc.RoleService.CreateRole(role.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating role"})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func (rc *RoleController) AssignRoleToProduct(c *gin.Context) {
	var productRole models.ProductRole
	if err := c.ShouldBindJSON(&productRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := rc.RoleService.AssignRoleToProduct(productRole.ProductID, productRole.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error assigning role to product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Role assigned to product successfully"})
}
