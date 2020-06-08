package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"../../utils"
	"log"
)

type StateSevc struct {
	db *repository.Repo
}


func NewStateService(db *repository.Repo) service.StateService {
	return &StateSevc{db:db}
}

func (r StateSevc) FindAll() ([]models.StateDTO, error) {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var states []models.State
	var statesDTO []models.StateDTO
	if dbErr := db.Find(&states).Error; dbErr != nil {
		return statesDTO,dbErr
	}
	for i,_ := range states {
		stateDto := utils.ParseStateToStateDTO(&states[i])
		statesDTO = append(statesDTO,stateDto)
	}
	return statesDTO,nil
}
