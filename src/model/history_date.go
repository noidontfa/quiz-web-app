package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type HistoryDate struct {
	gorm.Model
	InDate 		time.Time `gorm:"not null"`
	Histories	[]History `gorm:"foreignkey:HistoryDateId`
}
