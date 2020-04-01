package service

import (
	models "../model"
)

type CategoryService interface {
	FindAll() []models.Category
	Save(category models.Category) models.Category
}
