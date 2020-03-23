package service

import (
	"../entity"
)

type PostService interface {
	Create(post *entity.Post ) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}