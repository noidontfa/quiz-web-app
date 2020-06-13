package utils

import (
	models "../model"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func ParseQuizToQuizDTO(quiz  *models.Quiz) models.QuizDTO {
	var totalRating float32 = 0
	for _, e := range quiz.Ratings {
		totalRating += e.Star
	}
	if len(quiz.Ratings) != 0 {
		totalRating = totalRating / float32(len(quiz.Ratings))
	}

	var questions []models.QuestionDTO
	for _,e := range quiz.Questions {
		questions = append(questions, ParseQuestionTOQuestionDTO(&e))
	}
	timeFormater := "2006/01/02"
	return models.QuizDTO{
		ID:            	quiz.ID,
		CreatedAt:     	quiz.CreatedAt.Format(timeFormater),
		Name:          	quiz.Name,
		Description:   	quiz.Description,
		CategoryRefer: 	ParseCategoryToCategoryDTO(&quiz.CategoryRefer),
		LanguageRefer: 	ParseLanguageToLanguageDTO(&quiz.LanguageRefer),
		TimingRefer:   	ParseTimingToTimingDTO(&quiz.TimingRefer),
		UserRefer:     	ParseUserToUserDTO(&quiz.UserRefer),
		StateRefer:		ParseStateToStateDTO(&quiz.StateRefer),
		TotalQuestions: len(questions),
		QuestionRefer: 	questions,
		Ratings:       	totalRating,
		Image:         	quiz.FileName,
	}

}

func ParseQuestionTOQuestionDTO(question *models.Question) models.QuestionDTO {
	var choices []models.ChoiceDTO
	for _, e := range question.Choices {
		choices = append(choices,ParseChoiceToChoiceDTO(&e))
	}

	return models.QuestionDTO{
		ID:      question.ID,
		Name:    question.Name,
		Choices: choices,
		Image:   question.Image,
	}
}

func ParseChoiceToChoiceDTO(choice *models.Choice) models.ChoiceDTO {
	return models.ChoiceDTO{
		ID:      choice.ID,
		Name:    choice.Name,
		IsRight: choice.IsRight,
		Image:   choice.Image,
	}
}

func ParseCategoryToCategoryDTO(category *models.Category) models.CategoryDTO {
	return models.CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ParseLanguageToLanguageDTO(language *models.Language) models.LanguageDTO  {
	return models.LanguageDTO{
		ID:   language.ID,
		Name: language.Name,
	}
}

func ParseTimingToTimingDTO(timing *models.Timing) models.TimingDTO {
	return models.TimingDTO{
		ID:   timing.ID,
		Name: timing.Name,
		Sec:  timing.Sec,
	}
}

func ParseUserToUserDTO(user *models.User) models.UserDTO {
	return models.UserDTO{
		ID:         user.ID,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		DayOfBirth: user.DayOfBirth,
	}
}

func ParseHistoryToHistoryDTO(history *models.History) models.HistoryDTO {
	return models.HistoryDTO{
		ID:                 history.ID,
		NumberRightAnswers: history.NumberRightAnswers,
		Score:              history.Score,
		QuizRefer:          ParseQuizToQuizDTO(&history.QuizRefer),
		UserRefer:          ParseUserToUserDTO(&history.UserRefer),
		CreateAt:			history.CreatedAt.Format("2006-01-02"),
	}
}

func ParseStateToStateDTO(state *models.State) models.StateDTO {
	return models.StateDTO{
		ID:   state.ID,
		Name: state.Name,
	}
}

func Random_filename_16_char() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}

func SaveImage(filename string, dataBase64 string) (string, error) {
	randFileName,_ := Random_filename_16_char()
	randFileName += "-" + filename
	dec, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		return "",err
	}
	f, err := os.Create("./src/public/" + randFileName)
	if err != nil {
		return "",err
	}
	if _, err := f.Write(dec); err != nil {
		return  "",err
	}
	if err := f.Sync(); err != nil {
		return  "",err
	}
	returnString := "/file/" + randFileName
	return returnString, nil
}
