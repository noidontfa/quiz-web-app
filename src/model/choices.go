package model

import "github.com/jinzhu/gorm"

type Choice struct {
	gorm.Model
	Name 		string `gorm:"not null"`
	QuestionId	uint
	IsRight		bool `gorm:"not null;default: false"`
}
