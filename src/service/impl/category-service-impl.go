package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"../../utils"
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

func (s *Sevc) FindAll() ([]models.CategoryDTO, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var categories []models.Category
	var categoriesDto []models.CategoryDTO
	if dbs := db.Find(&categories).Error; dbs != nil {
		return categoriesDto, dbs
	}
	for i, _ := range categories {
		categoryDto := utils.ParseCategoryToCategoryDTO(&categories[i])
		categoriesDto = append(categoriesDto, categoryDto)
	}
	return categoriesDto, nil
}

func (s *Sevc) Save(category *models.Category) (*models.Category, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbs := db.Save(category).Error; dbs == nil {
		return category, nil
	} else {
		return &models.Category{}, dbs
	}
}

func (s *Sevc) FindById(id uint) (*models.Category, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var category models.Category
	if dbErr := db.Where("id = ?", id).Preload("Quizzes").Find(&category).Error; dbErr == nil {
		//db.Debug().Model(category).Association("quizzes").Find(&category.Quizzes)
		for i, _ := range category.Quizzes {
			quiz := &category.Quizzes[i]
			db.Model(quiz).Related(&quiz.CategoryRefer)
			db.Model(quiz).Related(&quiz.TimingRefer)
			db.Model(quiz).Related(&quiz.LanguageRefer)
			db.Model(quiz).Related(&quiz.UserRefer, "createdBy")
		}
		return &category, nil
	} else {
		return &models.Category{}, dbErr
	}
}

func (s *Sevc) Update(id uint, category *models.Category) (*models.Category, error) {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	if dbErr := db.Model(models.Category{}).Where("id = ?", id).Update(category).Find(category).Error; dbErr == nil {
		return category, nil
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
