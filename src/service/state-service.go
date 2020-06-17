package service

import (
	models "../model"
)

type StateService interface {
	FindAll() ([]models.StateDTO, error)
}
