package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pesto_coding/user_service/app/controllers"
	"github.com/pesto_coding/user_service/app/models"
	"github.com/pesto_coding/user_service/app/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoleController(t *testing.T) {
	roleService := setupTestRoleService()
	roleController := controllers.NewRoleController(roleService)
	router := gin.Default()

	router.GET("/roles/:name", roleController.GetRole)
	router.POST("/roles", roleController.CreateRole)
	router.POST("/roles/assign", roleController.AssignRoleToUser)

	// Test CreateRole
	req, _ := http.NewRequest("POST", "/roles", strings.NewReader(`{"name":"admin"}`))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test GetRole
	req, _ = http.NewRequest("GET", "/roles/admin", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUserController(t *testing.T) {
	userService := setupTestUserService()
	userController := controllers.NewUserController(userService)
	router := gin.Default()

	router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
	router.POST("/users/authenticate", userController.AuthenticateUser)

	// Test CreateUser
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(`{"username":"testuser", "email":"testuser@gmail.com", "password":"password"}`))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test GetUser
	id := "1"
	req, _ = http.NewRequest("GET", "/users/"+id, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func setupTestUserService() *services.UserService {
	db := setupTestDB()
	userStore := models.NewUserStore(db)
	return services.NewUserService(userStore)
}

func setupTestRoleService() *services.RoleService {
	db := setupTestDB()
	roleStore := models.NewRoleStore(db)
	return services.NewRoleService(roleStore)
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Role{})

	return db
}
