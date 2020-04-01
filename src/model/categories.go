package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name 		string `gorm:"not null"`
	Quizzes		[]Quiz `gorm:"foreignkey:CategoryId`
}
