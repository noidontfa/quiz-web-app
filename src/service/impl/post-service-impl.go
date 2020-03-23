package impl

import (
	"../../entity"
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


func (s *serv) Create(post *entity.Post) (*entity.Post, error) {
	return s.repo.Save(post)
}

func (s *serv) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}