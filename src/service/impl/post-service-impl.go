package impl

import (
	"../../model"
	"../../repository"
	"../../repository/impl"
	"../../service"
)

type serv struct {
	repo repository.PostRepository
}


func NewPostService() service.PostService {
	return &serv{
		repo: impl.NewPostRepository(),
	}
}


func (s *serv) Create(post *model.Post) (*model.Post, error) {
	return s.repo.Save(post)
}

func (s *serv) FindAll() ([]model.Post, error) {
	return s.repo.FindAll()
}