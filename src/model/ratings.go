package model

import "time"

type Rating struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time 	`json:"-"`
	UpdatedAt 	time.Time 	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
	UserId		uint 		`json:"userId"`
	QuizId 		uint 		`json:"quizId"`
	Star 		uint8 		`json:"star"`
}
