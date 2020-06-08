package service

import (
	models "../model"
)

type LanguageService interface {
	FindAll() ([]models.LanguageDTO, error)
	FindById(id uint) (*models.Language, error)
	Save(language *models.Language) (*models.Language, error)
	Update(id uint,language *models.Language) (*models.Language, error)
	Delete(id uint) error
}
