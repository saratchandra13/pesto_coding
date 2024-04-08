package models

import "github.com/jinzhu/gorm"

type RoleStore struct {
	db *gorm.DB
}

func NewRoleStore(db *gorm.DB) *RoleStore {
	return &RoleStore{
		db: db,
	}
}

func (rs *RoleStore) GetRoleByName(name string) *Role {
	var role Role
	rs.db.Where("name = ?", name).First(&role)
	return &role
}

func (rs *RoleStore) GetRoleById(id int) *Role {
	var role Role
	rs.db.Find(&role, id)
	return &role
}

func (rs *RoleStore) CreateRole(name string) (*Role, error) {
	role := &Role{Name: name}
	err := rs.db.Create(role).Error
	return role, err
}

func (rs *RoleStore) AssignRoleToUser(user *User, role *Role) error {
	user.Roles = append(user.Roles, *role)
	return rs.db.Save(user).Error
}
