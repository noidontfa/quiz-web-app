package model

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Name 		string
	QuizId		uint
	Choices		[]Choice `gorm:"foreignkey:QuestionId`
}
