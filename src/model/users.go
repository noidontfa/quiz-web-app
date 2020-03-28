package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username 	string
	Password 	string
	FirstName	string
	LastName	string
	Email		string
	DayOfBirth  time.Time
	Roles		[]Role `gorm:"many2many:user_roles"`
	Quizzes		[]Quiz `gorm:"foreignkey:CreatedBy"`
	Ratings		[]Rating `gorm:"foreignkey:UserId"`
	Histories	[]History `gorm:"foreignkey:UserId"`
}
