package model

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name	string
	Quizzes []Quiz `gorm:"foreignkey:LanguageId`
}
