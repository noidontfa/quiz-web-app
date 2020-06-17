package model

import (
	"time"
)

type Timing struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `gorm:"not null" json:"name"`
	Sec       uint8      `gorm:"not null" json:"sec"`
	Quizzes   []Quiz     `gorm:"foreignkey:TimingId" json:"quizzes"`
}

type TimingDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Sec  uint8  `json:"sec"`
}
