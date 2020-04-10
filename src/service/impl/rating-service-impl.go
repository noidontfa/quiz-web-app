package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type RatingServ struct {
	db *repository.Repo
}

func NewRatingService(db *repository.Repo) service.RatingService {
	return &RatingServ{db:db}
}

func (r *RatingServ) Save(rating *models.Rating) (*models.Rating, error) {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if rating.ID > 0 {
		dbErr := db.Model(rating).Update(rating).Error
		return rating,dbErr
	} else {
		dbErr := db.Save(rating).Error
		return rating,dbErr
	}
}


