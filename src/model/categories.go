package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name 		string
	Quizzes		[]Quiz `gorm:"foreignkey:CategoryId`
}
