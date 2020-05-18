package model

import (
	"time"
)

type Choice struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"deleteAt"`
	Name 		string 		`json:"name" sql:"not null"`
	QuestionId	uint		`json:"questionId"`
	IsRight		bool 		`gorm:"default: false" json:"isRight"`
	Image		string		`json:"image"`
}

type ChoiceDTO struct {
	ID        	uint 		`json:"id"`
	Name 		string 		`json:"name"`
	IsRight		bool 		`json:"isRight"`
	Image		string		`json:"image"`
}