package service

import (
	models "../model"
)

type RoleService interface {
	FindAll() ([]models.Role, error)
	FindById(id uint) (*models.Role, error)
	Save(role *models.Role) (*models.Role, error)
	Update(id uint,role *models.Role) (*models.Role, error)
	Delete(id uint) error
}
