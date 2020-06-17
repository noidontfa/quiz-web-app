package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type LoginSevc struct {
	db *repository.Repo
}

func NewLoginService(db *repository.Repo) service.LoginService {
	return &LoginSevc{db: db}
}

func (l *LoginSevc) Login(username string, password string) bool {
	db, err := l.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var user models.User
	db.Where("username = ? AND password = ? ", username, password).Find(&user)
	return user.ID > 0
}
