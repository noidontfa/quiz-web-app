package utils

import (
	models "../model"
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

	return models.QuizDTO{
		ID:            quiz.ID,
		CreatedAt:     quiz.CreatedAt,
		Name:          quiz.Name,
		Description:   quiz.Description,
		CategoryRefer: ParseCategoryToCategoryDTO(&quiz.CategoryRefer),
		LanguageRefer: ParseLanguageToLanguageDTO(&quiz.LanguageRefer),
		TimingRefer:   ParseTimingToTimingDTO(&quiz.TimingRefer),
		UserRefer:     ParseUserToUserDTO(&quiz.UserRefer),
		TotalQuestions: len(questions),
		QuestionRefer: questions,
		Ratings:       totalRating,
		Image:         quiz.Image,
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