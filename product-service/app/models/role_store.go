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

func (rs *RoleStore) AssignRoleToProduct(productId uint, roleId uint) error {
	product := rs.db.Find(&Product{}, productId).Value.(*Product)
	role := rs.db.Find(&Role{}, roleId).Value.(*Role)
	rs.db.Model(product).Association("Roles").Append(role)
	return nil
}
