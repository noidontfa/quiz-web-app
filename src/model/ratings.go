package model

import "github.com/jinzhu/gorm"

type Rating struct {
	gorm.Model
	UserId	uint
	QuizId 	uint
	Star 	uint8
}
