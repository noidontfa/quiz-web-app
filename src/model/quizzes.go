package model

import (
	"time"
)

type Quiz struct {
	ID        		uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 		time.Time 	`json:"createdAt"`
	UpdatedAt 		time.Time 	`json:"updatedAt"`
	DeletedAt 		*time.Time 	`sql:"index" json:"-"`
	Name			string 		`gorm:"not null" json:"name"`
	Description 	string 		`json:"description"`
	CreatedBy		uint 		`gorm:"not null" json:"createdBy"`
	CategoryId		uint 		`json:"categoryId"`
	LanguageId  	uint 		`json:"languageId"`
	TimingId    	uint 		`json:"timingId"`
	StateId			uint		`json:"stateId"`
	CategoryRefer 	Category 	`json:"categoryRefer"`
	LanguageRefer 	Language	`json:"languageRefer"`
	TimingRefer		Timing		`json:"timingRefer"`
	StateRefer		State		`json:"stateRefer"`
	UserRefer		User		`json:"userRefer"`
	Questions		[]Question 	`gorm:"foreignkey:QuizId" json:"questions"`
	Ratings			[]Rating 	`gorm:"foreignkey:QuizId" json:"ratings"`
	Histories		[]History 	`gorm:"foreignkey:QuizId" json:"histories"`
	Image			string		`json:"image" gorm:"-"`
	FileName		string		`json:"filename"`
}

type QuizDTO struct {
	ID        		uint 			`json:"id"`
	CreatedAt 		string 		`json:"createdAt"`
	Name			string 			`json:"name"`
	Description 	string 			`json:"description"`
	CategoryRefer 	CategoryDTO 	`json:"categoryRefer"`
	LanguageRefer 	LanguageDTO		`json:"languageRefer"`
	TimingRefer		TimingDTO		`json:"timingRefer"`
	UserRefer		UserDTO			`json:"userRefer"`
	StateRefer		StateDTO		`json:"stateRefer"`
	TotalQuestions	int				`json:"totalQuestions"`
	QuestionRefer	[]QuestionDTO 	`json:"questionRefer"`
	Ratings			float32			`json:"ratings"`
	Image			string			`json:"image"`
}
