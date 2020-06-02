package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"../../utils"
	"fmt"
	"log"
)

type HistoryServ struct {
	db *repository.Repo
}

func NewHistoryService(db *repository.Repo) service.HistoryService {
	return &HistoryServ{db:db}
}

func (h *HistoryServ) FindByQuizId(id uint) ([]models.HistoryDTO, error) {
	db, err := h.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	var historiesDTO []models.HistoryDTO
	var histories []models.History
	dbErr := db.Where("quiz_id = ?", id).Find(&histories).Error
	if dbErr == nil {
		for i,_ := range histories {
			history := &histories[i]
			db.Model(history).Related(&history.UserRefer)
			db.Model(history).Related(&history.QuizRefer)
			historiesDTO = append(historiesDTO,utils.ParseHistoryToHistoryDTO(history))
		}
	}
	return historiesDTO,dbErr
}

func (h *HistoryServ) FindByUserId(id uint) ([]models.HistoryDTO, error) {
	db, err := h.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	var historiesDTO []models.HistoryDTO
	var histories []models.History
	dbErr := db.Where("user_id = ?", id).Find(&histories).Error
	if dbErr == nil {
		for i,_ := range histories {
			history := &histories[i]
			db.Model(history).Related(&history.UserRefer)
			db.Model(history).Related(&history.QuizRefer)
			historiesDTO = append(historiesDTO,utils.ParseHistoryToHistoryDTO(history))
		}
	}

	return historiesDTO,dbErr
}


func (h *HistoryServ) Save(history *models.History) (*models.HistoryDTO, error) {
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

	historyDTO := utils.ParseHistoryToHistoryDTO(history)

	return &historyDTO,dbErr
}


func (h *HistoryServ) FindByDateId(date string, quizId uint) ([]models.HistoryDTO, error) {
	db, err := h.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var histories []models.History
	dataf := fmt.Sprintf("quiz_id = %d AND created_at LIKE '%s'", quizId, date + "%")
	dbErr := db.Where(dataf).Find(&histories).Error
	if dbErr != nil {
		log.Fatal(err.Error())
	}
	var historiesDTO []models.HistoryDTO
	for i, _ := range histories {
		history := &histories[i]
		db.Model(history).Related(&history.UserRefer)
		db.Model(history).Related(&history.QuizRefer)
		historiesDTO = append(historiesDTO, utils.ParseHistoryToHistoryDTO(history))
	}
	return historiesDTO,err
}


