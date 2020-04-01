package model

import "github.com/jinzhu/gorm"

type Timing struct {
	gorm.Model
	Name	string	`gorm:"not null"`
	Sec 	string	`gorm:"not null"`
	Quizzes	[]Quiz 	`gorm:"foreignkey:TimingId`
}
