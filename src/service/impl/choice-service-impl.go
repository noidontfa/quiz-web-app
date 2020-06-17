package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type CServ struct {
	db *repository.Repo
}

func NewChoiceService(db *repository.Repo) service.ChoiceService {
	return &CServ{db: db}
}

func (c *CServ) Delete(id uint) error {
	db, err := c.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	dbErr := db.Where("id = ?", id).Delete(&models.Choice{}).Error

	return dbErr
}
