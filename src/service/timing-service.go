package service

import (
	models "../model"
)

type TimingService interface {
	FindAll() ([]models.Timing, error)
	FindById(id uint) (*models.Timing, error)
	Save(timing *models.Timing) (*models.Timing, error)
	Update(id uint,timing *models.Timing) (*models.Timing, error)
	Delete(id uint) error
}
