package service

import (
	models "../model"
)

type RatingService interface {
	Save(rating *models.Rating) (*models.Rating, error)
}
