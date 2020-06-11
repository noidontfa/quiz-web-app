package service

import (
	models "../model"
)

type QuestionService interface {
	Save(quizId uint,questions []models.Question) ([]models.QuestionDTO,error)
	Delete(id uint) error
}