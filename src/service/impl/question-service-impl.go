package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type QuestionServ struct {
	db *repository.Repo
}

func NewQuestionService(db *repository.Repo) service.QuestionService {
	return &QuestionServ{db:db}
}

func (q *QuestionServ) Save(quizId uint, questions []models.Question) ([]models.Question, error) {
	db, err := q.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx := db.Begin() // open transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return questions,err
	}

	for i, _ := range questions {
		question := &questions[i]
		questionId := question.ID
		question.QuizId = quizId
		choices := question.Choices
		question.Choices = nil
		if questionId > 0 {
			dbErr := tx.Model(question).Update(question).Error
			if dbErr != nil {
				tx.Rollback()
				return questions,dbErr
			}
		} else {
			dbErr := tx.Save(&question).Error
			if dbErr != nil {
				tx.Rollback()
				return questions,dbErr
			}
		}
		questionId = question.ID
		for j, _ := range choices {
			choice := &choices[j]
			choice.QuestionId = questionId
			choiceId := choice.ID
			if choiceId > 0 {
				dbErr := tx.Model(choice).Update(choice).Error
				if dbErr != nil {
					tx.Rollback()
					return questions,dbErr
				}
			} else {
				dbErr := tx.Save(&choice).Error
				if dbErr != nil {
					tx.Rollback()
					return questions,dbErr
				}
			}
		}
		dbErr := tx.Model(question).Related(&question.Choices).Error
		if dbErr != nil {
			tx.Rollback()
			return questions,dbErr
		}
	}

	return questions, tx.Commit().Error
}

func (q *QuestionServ) Delete(id uint) error {
	db, err := q.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	dbErr := db.Where("id = ?", id).Delete(&models.Question{}).Error

	return dbErr
}


