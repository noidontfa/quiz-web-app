package service

import (
	models "../model"
)

type QuizService interface {
	FindAll() ([]models.QuizDTO, error)
	FindById(id uint) (*models.QuizDTO, error)
	Save(quiz *models.Quiz) (*models.QuizDTO, error)
	Update(id uint,quiz *models.Quiz) (*models.QuizDTO, error)
	Delete(id uint) error
}
