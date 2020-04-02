package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type Sevc struct {
	db *repository.Repo
}

func NewCategoryService(db *repository.Repo) service.CategoryService {
	return &Sevc{
		db: db,
	}
}

func (s *Sevc) FindAll() ([]models.Category, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var categories []models.Category
	if dbs := db.Find(&categories).Error; dbs == nil {
		return categories,nil
	} else {
		return []models.Category{}, dbs
	}

}

func (s *Sevc) Save(category *models.Category) (*models.Category,error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbs := db.Save(category).Error; dbs == nil {
		return category,nil
	} else {
		return &models.Category{},dbs
	}
}


func (s *Sevc) FindById(id uint) (*models.Category, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var category models.Category
	if dbErr := db.Where("id = ?", id).Find(&category).Error; dbErr == nil {
		return &category,nil
	} else {
		return &models.Category{},dbErr
	}
}


func (s *Sevc) Update(id uint,category *models.Category) (*models.Category, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Model(models.Category{}).Where("id = ?", id).Update(category).Find(category).Error; dbErr == nil {
		return category,nil
	} else {
		return &models.Category{}, dbErr
	}
}

func (s *Sevc) Delete(id uint) error {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Where("id = ?", id).Delete(models.Category{}).Error; dbErr == nil {
		return nil
	} else {
		return dbErr
	}
}



