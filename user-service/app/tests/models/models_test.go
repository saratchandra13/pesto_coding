package models_test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pesto_coding/user_service/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoleStore(t *testing.T) {
	db := setupTestDB()
	roleStore := models.NewRoleStore(db)

	// Test CreateRole
	role, err := roleStore.CreateRole("admin")
	assert.Nil(t, err)
	assert.Equal(t, "admin", role.Name)

	// Test GetRoleByName
	role = roleStore.GetRoleByName("admin")
	assert.Nil(t, err)
	assert.Equal(t, "admin", role.Name)
}

func TestUserStore(t *testing.T) {
	db := setupTestDB()
	userStore := models.NewUserStore(db)

	// Test CreateUser
	user, err := userStore.CreateUser("testuser", "testpassword@gmail.com", "password")
	assert.Nil(t, err)
	assert.Equal(t, "testuser", user.Username)

	// Test GetUserByUsername
	user = userStore.GetUserByUsername("testuser")
	assert.Equal(t, "testuser", user.Username)
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
