package service

import (
	models "../model"
)

type QuizService interface {
	FindAll() ([]models.Quiz, error)
	FindById(id uint) (*models.Quiz, error)
	Save(quiz *models.Quiz) (*models.Quiz, error)
	Update(id uint,quiz *models.Quiz) (*models.Quiz, error)
	Delete(id uint) error
}
