package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)
type QuizSevc struct {
	db *repository.Repo
}


func NewQuizService(db *repository.Repo) service.QuizService {
	return &QuizSevc{
		db: db,
	}
}

func (q *QuizSevc) FindAll() ([]models.Quiz, error) {
	db, err := q.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var quizzes []models.Quiz
	if dbErr := db.Find(&quizzes).Error; dbErr == nil {
		for i,_ := range quizzes {
			db.Model(quizzes[i]).Related(&quizzes[i].CategoryRefer)
			db.Model(quizzes[i]).Related(&quizzes[i].LanguageRefer)
			db.Model(quizzes[i]).Related(&quizzes[i].TimingRefer)
			db.Model(quizzes[i]).Related(&quizzes[i].UserRefer,"CreatedBy")
		}
		return quizzes,nil
	} else {
		return []models.Quiz{},dbErr
	}
}

func (q *QuizSevc) FindById(id uint) (*models.Quiz, error) {
	db, err := q.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var quiz models.Quiz
	dbErr := db.Where("id = ?", id).Find(&quiz).Error
	if dbErr == nil {
		db.Model(quiz).Related(&quiz.CategoryRefer)
		db.Model(quiz).Related(&quiz.LanguageRefer)
		db.Model(quiz).Related(&quiz.TimingRefer)
		db.Model(quiz).Related(&quiz.UserRefer,"CreatedBy")
		if dbErr := db.Model(quiz).Association("Questions").Find(&quiz.Questions).Error; dbErr == nil {
			for i,_ := range quiz.Questions {
				question := &quiz.Questions[i]
				db.Model(question).Association("Choices").Find(&question.Choices)
			}
		}
	}
	return &quiz, dbErr
}

func (q *QuizSevc) Save(quiz *models.Quiz) (*models.Quiz, error) {
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
		return quiz,err
	}

	dbErr := tx.Save(quiz).Error
	if dbErr == nil {
		quizId := quiz.ID
		for i, _ := range quiz.Questions {
			question := &quiz.Questions[i]
			question.QuizId = quizId
			if dbErr := tx.Save(question).Error; dbErr == nil {
				questionId := question.ID
				for j, _ := range question.Choices {
					choice := &question.Choices[j]
					choice.QuestionId = questionId
					if dbErr := tx.Save(choice).Error; dbErr != nil {
						tx.Rollback()
						return quiz, dbErr
					}
				}
			} else {
				tx.Rollback()
				return quiz,dbErr
			}
		}
		db.Model(quiz).Related(&quiz.CategoryRefer)
		db.Model(quiz).Related(&quiz.LanguageRefer)
		db.Model(quiz).Related(&quiz.TimingRefer)
		db.Model(quiz).Related(&quiz.UserRefer,"CreatedBy")
	} else {
		tx.Rollback()
		return quiz,dbErr
	}

	return quiz, tx.Commit().Error
}

func (q *QuizSevc) Update(id uint, quiz *models.Quiz) (*models.Quiz, error) {
	db, err := q.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	dbErr := db.Model(&models.Quiz{}).Where("id = ?", id).Update(quiz).Find(quiz).Error
	if dbErr == nil {
		db.Model(quiz).Related(&quiz.CategoryRefer)
		db.Model(quiz).Related(&quiz.LanguageRefer)
		db.Model(quiz).Related(&quiz.TimingRefer)
		db.Model(quiz).Related(&quiz.UserRefer, "CreatedBy")
		if dbErr := db.Model(quiz).Association("Questions").Find(&quiz.Questions).Error; dbErr == nil {
			for i, _ := range quiz.Questions {
				question := &quiz.Questions[i]
				db.Model(question).Association("Choices").Find(&question.Choices)
			}
		}

	}
	return quiz,dbErr
}

func (q *QuizSevc) Delete(id uint) error {
	db, err := q.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Where("id = ?", id).Delete(&models.Quiz{}).Error; dbErr == nil {
		return nil
	} else {
		return dbErr
	}
}


