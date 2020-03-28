package model

import "github.com/jinzhu/gorm"

type Timing struct {
	gorm.Model
	Name	string
	Sec 	string
	Quizzes	[]Quiz `gorm:"foreignkey:TimingId`
}
