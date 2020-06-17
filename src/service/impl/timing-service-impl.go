package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"../../utils"
	"log"
)

type TimingSevc struct {
	db *repository.Repo
}

func NewTimingService(db *repository.Repo) service.TimingService {
	return &TimingSevc{db: db}
}

func (t *TimingSevc) FindAll() ([]models.TimingDTO, error) {
	db, err := t.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var timings []models.Timing
	var timingsDto []models.TimingDTO
	if dbErr := db.Find(&timings).Error; dbErr != nil {
		return timingsDto, dbErr
	}
	for i, _ := range timings {
		timingDto := utils.ParseTimingToTimingDTO(&timings[i])
		timingsDto = append(timingsDto, timingDto)
	}
	return timingsDto, nil
}

func (t *TimingSevc) FindById(id uint) (*models.Timing, error) {
	db, err := t.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var timing models.Timing
	if dbErr := db.Where("id = ?", id).Preload("Quizzes").Find(&timing).Error; dbErr == nil {
		for i, _ := range timing.Quizzes {
			quiz := &timing.Quizzes[i]
			db.Model(quiz).Related(&quiz.CategoryRefer)
			db.Model(quiz).Related(&quiz.TimingRefer)
			db.Model(quiz).Related(&quiz.LanguageRefer)
			db.Model(quiz).Related(&quiz.UserRefer, "createdBy")
		}
		return &timing, nil
	} else {
		return &timing, dbErr
	}
}

func (t *TimingSevc) Save(timing *models.Timing) (*models.Timing, error) {
	db, err := t.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Save(timing).Error; dbErr == nil {
		return timing, nil
	} else {
		return timing, dbErr
	}
}

func (t *TimingSevc) Update(id uint, timing *models.Timing) (*models.Timing, error) {
	db, err := t.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Model(timing).Where("id = ?", id).Update(&timing).Find(&timing).Error; dbErr == nil {
		return timing, nil
	} else {
		return timing, dbErr
	}

}

func (t *TimingSevc) Delete(id uint) error {
	db, err := t.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Where("id = ?", id).Delete(&models.Timing{}).Error; dbErr == nil {
		return nil
	} else {
		return dbErr
	}
}
