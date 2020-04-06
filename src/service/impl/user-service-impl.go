package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type UserServ struct {
	db *repository.Repo
}

func NewUserService(db *repository.Repo) service.UserService  {
	return &UserServ{db:db}
}

func (u *UserServ) FindAll() ([]models.User, error) {
	db, err := u.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var users []models.User
	dbErr := db.Find(&users).Error
	if dbErr == nil {
		for i,_ := range users {
			user := &users[i]
			db.Model(user).Association("Roles").Find(&user.Roles)
		}
	}
	return users,dbErr
}

func (u *UserServ) FindById(id uint) (*models.User, error) {
	panic("implement me")
}

func (u *UserServ) Save(user *models.User) (*models.User, error) {
	db, err := u.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var roles []models.Role
	if dbErr := db.Where(user.RoleIds).Find(&roles).Error; dbErr == nil {
		dbErr1 := db.Debug().Save(user).Association("Roles").Append(roles).Error
		return user,dbErr1
	} else {
		return &models.User{},nil
	}
}

func (u *UserServ) Update(id uint, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u *UserServ) Delete(id uint) error {
	panic("implement me")
}
