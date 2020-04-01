package service

import (
	models "../model"
)

type CategoryService interface {
	FindAll() ([]models.Category, error)
	FindById(id uint) (models.Category, error)
	Save(category models.Category) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(id uint) error
}
