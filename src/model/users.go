package model

import (
	"time"
)

type User struct {
	ID        	uint 		`gorm:"primary_key"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
	Username 	string		`gorm:"not null;unique" json:"username"`
	Password 	string		`gorm:"not null" json:"password"`
	FirstName	string		`gorm:"not null" json:"firstName"`
	LastName	string		`gorm:"not null" json:"lastName"`
	Email		string		`gorm:"not null" json:"email"`
	DayOfBirth  time.Time	`json:"dayOfBirth" time_format:"2006-01-02"`
	RoleIds		[]uint		`gorm:"-" json:"roleIds"`
	Roles		[]Role 		`gorm:"many2many:user_roles" json:"roles"`
	Quizzes		[]Quiz 		`gorm:"foreignkey:CreatedBy;" json:"quizzes"`
	Ratings		[]Rating 	`gorm:"foreignkey:UserId" json:"ratings"`
	Histories	[]History 	`gorm:"foreignkey:UserId" json:"histories"`
}