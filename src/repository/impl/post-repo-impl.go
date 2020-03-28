package impl

import (
	"../../model"
	"../../repository"
)


type repo struct {
}

func NewPostRepository() repository.PostRepository {
	return &repo{}
}

func (r *repo)Save(post *model.Post) (*model.Post,error) {
	panic("Me Save")
}

func (r *repo) FindAll() ([]model.Post, error) {
	panic("Me Find All")
}