package model

import (
	"time"
)

type Role struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt	*time.Time 	`sql:"index" json:"-"`
	Name 		string 		`gorm:"not null" json:"name"`
	Users		[]User		`gorm:"many2many:user_roles" json:"users"`
}