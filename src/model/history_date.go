package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type HistoryDate struct {
	gorm.Model
	InDate 		time.Time
	Histories	[]History `gorm:"foreignkey:HistoryDateId`
}
