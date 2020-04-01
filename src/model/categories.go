package model

import (
	"time"
)

type Category struct {
	ID        	uint `gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
	DeletedAt 	*time.Time `sql:"index" json:"deleteAt"`
	Name 		string `gorm:"not null" json:"name"`
	Quizzes		[]Quiz `gorm:"foreignkey:CategoryId" json:"-"`
}
