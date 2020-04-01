package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username 	string	`gorm:"not null;unique"`
	Password 	string	`gorm:"not null"`
	FirstName	string	`gorm:"not null"`
	LastName	string	`gorm:"not null"`
	Email		string	`gorm:"not null"`
	DayOfBirth  time.Time
	Roles		[]Role `gorm:"many2many:user_roles"`
	Quizzes		[]Quiz `gorm:"foreignkey:CreatedBy"`
	Ratings		[]Rating `gorm:"foreignkey:UserId"`
	Histories	[]History `gorm:"foreignkey:UserId"`
}
