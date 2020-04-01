package model

import "github.com/jinzhu/gorm"

type History struct {
	gorm.Model
	NumberRightAnswers int8 `gorm:"not null"`
	Score uint 				`gorm:"not null"`
	HistoryDateId uint
	QuizId uint
	UserId uint
}
