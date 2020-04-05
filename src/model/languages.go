package model

import (
	"time"
)

type Language struct {
	ID        		uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 		time.Time 	`json:"-"`
	UpdatedAt 		time.Time 	`json:"-"`
	DeletedAt 		*time.Time 	`sql:"index" json:"-"`
	Name			string 		`gorm:"not null" json:"name"`
	Quizzes 		[]Quiz 		`gorm:"foreignkey:LanguageId" json:"quizzes"`
}
