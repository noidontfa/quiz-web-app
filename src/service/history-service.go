package service

import (
	models "../model"
)

type HistoryService interface {
	FindByQuizId(id uint) ([]models.History, error)
	FindByUserId(id uint) ([]models.History, error)
	Save(history *models.History) (*models.History, error)
}
