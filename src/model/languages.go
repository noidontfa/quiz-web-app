package model

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name	string `gorm:"not null"`
	Quizzes []Quiz `gorm:"foreignkey:LanguageId`
}
