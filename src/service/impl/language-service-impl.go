package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type LanguageSevc struct {
	db *repository.Repo
}

func NewLanguageService(db *repository.Repo) service.LanguageService {
	return &LanguageSevc{db:db}
}

func (l *LanguageSevc) FindAll() ([]models.Language, error) {
	db, err := l.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var languages []models.Language
	if dbErr := db.Find(&languages).Error; dbErr == nil {
		return languages,nil
	} else {
		return languages,dbErr
	}
}

func (l *LanguageSevc) FindById(id uint) (*models.Language, error) {
	db, err := l.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var language models.Language
	if dbErr := db.Where("id = ?", id).Preload("Quizzes").Find(&language).Error; dbErr == nil {
		for i, _ := range language.Quizzes {
			quiz := &language.Quizzes[i]
			db.Model(quiz).Related(&quiz.CategoryRefer)
			db.Model(quiz).Related(&quiz.TimingRefer)
			db.Model(quiz).Related(&quiz.LanguageRefer)
			db.Model(quiz).Related(&quiz.UserRefer,"createdBy")
		}
		return &language,nil
	} else {
		return &language,dbErr
	}
}

func (l *LanguageSevc) Save(language *models.Language) (*models.Language, error) {
	db, err := l.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Save(language).Error; dbErr == nil {
		return language,nil
	} else {
		return language,dbErr
	}
}

func (l *LanguageSevc) Update(id uint, language *models.Language) (*models.Language, error) {
	db, err := l.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Model(language).Where("id = ?", id).Update(&language).Find(&language).Error; dbErr == nil {
		return language,nil
	} else {
		return language,dbErr
	}
}

func (l *LanguageSevc) Delete(id uint) error {
	db, err := l.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Where("id = ?", id).Delete(&models.Language{}).Error; dbErr == nil {
		return nil
	} else {
		return dbErr
	}
}

