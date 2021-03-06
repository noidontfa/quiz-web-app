package model

import "time"

type History struct {
	ID                 uint       `gorm:"primary_key" json:"id"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"-"`
	DeletedAt          *time.Time `sql:"index" json:"-"`
	NumberRightAnswers int8       `gorm:"not null" json:"numberRightAnswers"`
	Score              float32    `gorm:"not null" json:"score"`
	QuizId             uint       `json:"quizId"`
	UserId             uint       `json:"userId"`
	QuizRefer          Quiz       `json:"quizRefer"`
	UserRefer          User       `json:"userRefer"`
}

type HistoryDTO struct {
	ID                 uint    `json:"id"`
	NumberRightAnswers int8    `json:"numberRightAnswers"`
	Score              float32 `json:"score"`
	QuizRefer          QuizDTO `json:"quizRefer"`
	UserRefer          UserDTO `json:"userRefer"`
	CreateAt           string  `json:"createAt"`
}
