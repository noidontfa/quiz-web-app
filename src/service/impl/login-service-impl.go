package impl

import (
	"../../repository"
	"../../service"
)

type LoginSevc struct {
	db *repository.Repo
}

func NewLoginService(db *repository.Repo) service.LoginService {
	return &LoginSevc{db:db}
}

func (l *LoginSevc) Login(username string, password string) bool {
	return username == "admin" && password == "admin"
}
