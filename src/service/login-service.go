package service

type LoginService interface {
	Login(username string, password string) bool
}
