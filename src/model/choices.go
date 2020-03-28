package model

import "github.com/jinzhu/gorm"

type Choice struct {
	gorm.Model
	Name 		string
	QuestionId	uint
	IsRight		bool
}
