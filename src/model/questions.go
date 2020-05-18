package model

import "time"

type Question struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time 	`json:"-"`
	UpdatedAt 	time.Time 	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"deleteAt"`
	Name 		string 		`gorm:"not null" json:"name" sql:"not null"`
	QuizId		uint 		`json:"quizId"`
	Choices		[]Choice 	`gorm:"foreignkey:QuestionId" json:"choices"`
	Image			string		`json:"image"`
}

type QuestionDTO struct {
	ID        	uint 		`json:"id"`
	Name 		string 		`json:"name"`
	Choices		[]ChoiceDTO `json:"choices"`
	Image		string		`json:"image"`
}

