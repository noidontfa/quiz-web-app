package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name 	string 	`gorm:"not null"`
	Users	[]User	`gorm:"many2many:user_roles"`
}