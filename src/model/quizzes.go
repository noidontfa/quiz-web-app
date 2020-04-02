package model

import "time"

type Quiz struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`sql:"index" json:"deleteAt"`
	Name		string 		`gorm:"not null" json:"name"`
	Description string 		`json:"description"`
	CreatedBy	uint 		`gorm:"not null" json:"createdBy"`
	CategoryId	uint 		`json:"categoryId"`
	LanguageId  uint 		`json:"languageId"`
	TimingId    uint 		`json:"timingId"`
	Questions	[]Question 	`gorm:"foreignkey:QuizId" json:"-"`
	Ratings		[]Rating 	`gorm:"foreignkey:QuizId" json:"-"`
	Histories	[]History 	`gorm:"foreignkey:QuizId" json:"-"`
}
