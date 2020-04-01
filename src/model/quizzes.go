package model

import "github.com/jinzhu/gorm"

type Quiz struct {
	gorm.Model
	Name		string `gorm:"not null"`
	Description string
	CreatedBy	uint `gorm:"not null"`
	CategoryId	uint
	LanguageId  uint
	TimingId    uint
	Questions	[]Question `gorm:"foreignkey:QuizId"`
	Ratings		[]Rating `gorm:"foreignkey:QuizId"`
	Histories	[]History `gorm:"foreignkey:QuizId"`
}
