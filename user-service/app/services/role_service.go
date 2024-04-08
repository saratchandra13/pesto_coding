package services

import (
	"github.com/pesto_coding/user_service/app/models"
)

type RoleService struct {
	RoleStore *models.RoleStore
}

func NewRoleService(roleStore *models.RoleStore) *RoleService {
	return &RoleService{
		RoleStore: roleStore,
	}
}

func (rs *RoleService) GetRoleByName(name string) (*models.Role, error) {
	return rs.RoleStore.GetRoleByName(name), nil
}

func (rs *RoleService) GetRoleById(id int) (*models.Role, error) {
	return rs.RoleStore.GetRoleById(id), nil
}

func (rs *RoleService) CreateRole(name string) (*models.Role, error) {
	return rs.RoleStore.CreateRole(name)
}

func (rs *RoleService) AssignRoleToUser(user *models.User, role *models.Role) error {
	return rs.RoleStore.AssignRoleToUser(user, role)
}
