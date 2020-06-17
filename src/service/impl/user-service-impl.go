package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"../../utils"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type UserServ struct {
	db *repository.Repo
}

func NewUserService(db *repository.Repo) service.UserService {
	return &UserServ{db: db}
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
		for i, _ := range users {
			user := &users[i]
			db.Model(user).Association("Roles").Find(&user.Roles)
		}
	}
	return users, dbErr
}

func (u *UserServ) FindById(id uint) (*models.User, error) {
	db, err := u.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var user models.User
	dbErr := db.Where("id = ?", id).Find(&user).Error
	if dbErr == nil {
		db.Model(user).Association("Roles").Find(&user.Roles)
		//db.Model(user).Association("Quizzes").Find(&user.Quizzes)
	}
	return &user, dbErr
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
		return user, dbErr1
	} else {
		return &models.User{}, dbErr
	}
}

func (u *UserServ) Update(id uint, user *models.User) (*models.User, error) {
	db, err := u.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var roles []models.Role
	if len(user.RoleIds) > 0 {
		if dbErr := db.Where(user.RoleIds).Find(&roles).Error; dbErr != nil {
			return &models.User{}, dbErr
		}
	}
	dbErr1 := db.Model(user).Where("id = ?", id).Update(&user).Association("Roles").Replace(roles).Error
	return user, dbErr1
}

func (u *UserServ) Delete(id uint) error {
	db, err := u.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	dbErr := db.Where("id = ?", id).Delete(&models.User{}).Error
	return dbErr
}

func (u *UserServ) FindByUsername(username string) (*models.UserDTO, error) {
	db, err := u.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var user models.User
	var userDto models.UserDTO
	dbErr := db.Where("username = ?", username).Find(&user).Error
	if dbErr == nil {
		db.Model(user).Association("Roles").Find(&user.Roles)
		userDto = utils.ParseUserToUserDTO(&user)
		//db.Model(user).Association("Quizzes").Find(&user.Quizzes)
	}
	return &userDto, dbErr
}

func (u *UserServ) FindByToken(jwtToken string) (*models.UserDTO, error) {
	var claims models.JWTCustomClaims
	_, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("ThinhTest"), nil
	})
	if err != nil {
		return nil, err
	}
	user, _ := u.FindByUsername(claims.Username)
	return user, nil
}
