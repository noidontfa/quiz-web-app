package repository

import (
	"../model"
)

type PostRepository interface {
	Save(post *model.Post) (*model.Post,error)
	FindAll() ([]model.Post, error)
}

