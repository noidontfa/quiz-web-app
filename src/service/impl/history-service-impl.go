package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type HistoryServ struct {
	db *repository.Repo
}



func NewHistoryService(db *repository.Repo) service.HistoryService {
	return &HistoryServ{db:db}
}

func (h *HistoryServ) FindByQuizId(id uint) ([]models.History, error) {
	db, err := h.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	var histories []models.History
	dbErr := db.Where("quiz_id = ?", id).Find(&histories).Error
	if dbErr == nil {
		for i,_ := range histories {
			history := &histories[i]
			db.Model(history).Related(&history.UserRefer)
			db.Model(history).Related(&history.QuizRefer)
		}
	}
	return histories,dbErr
}

func (h *HistoryServ) FindByUserId(id uint) ([]models.History, error) {
	db, err := h.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	var histories []models.History
	dbErr := db.Where("user_id = ?", id).Find(&histories).Error
	if dbErr == nil {
		for i,_ := range histories {
			history := &histories[i]
			db.Model(history).Related(&history.UserRefer)
			db.Model(history).Related(&history.QuizRefer)
		}
	}
	return histories,dbErr
}


func (h *HistoryServ) Save(history *models.History) (*models.History, error) {
	db, err := h.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	dbErr := db.Save(history).Error
	if dbErr == nil {
		db.Model(history).Related(&history.UserRefer)
		db.Model(history).Related(&history.QuizRefer)
	}
	return history,dbErr
}


