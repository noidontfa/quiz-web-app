package service

import (
	"../model"
)

type PostService interface {
	Create(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}
