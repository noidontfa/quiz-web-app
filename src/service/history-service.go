package service

import (
	models "../model"
)

type HistoryService interface {
	FindByQuizId(id uint) ([]models.HistoryDTO, error)
	FindByUserId(id uint) ([]models.HistoryDTO, error)
	Save(history *models.History) (*models.HistoryDTO, error)
	FindByDateId(date string, quizId uint) ([]models.HistoryDTO,error)
}
