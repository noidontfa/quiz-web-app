package service

import (
	models "../model"
)

type UserService interface {
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	Save(user *models.User) (*models.User, error)
	Update(id uint,user *models.User) (*models.User, error)
	Delete(id uint) error
}
