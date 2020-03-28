package model

import "github.com/jinzhu/gorm"

type History struct {
	gorm.Model
	NumberRightAnswers int8
	Score uint
	HistoryDateId uint
	QuizId uint
	UserId uint
}
