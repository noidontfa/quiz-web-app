package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"log"
)

type RoleSevc struct {
	db *repository.Repo
}

func NewRoleService(db *repository.Repo) service.RoleService {
	return &RoleSevc{db: db}
}

func (r *RoleSevc) FindAll() ([]models.Role, error) {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var roles []models.Role
	if dbErr := db.Find(&roles).Error; dbErr == nil {
		return roles, nil
	} else {
		return roles, dbErr
	}
}

func (r *RoleSevc) FindById(id uint) (*models.Role, error) {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	var role models.Role
	dbErr := db.Where("id = ?", id).Find(&role).Error

	return &role, dbErr
}

func (r *RoleSevc) Save(role *models.Role) (*models.Role, error) {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	dbErr := db.Save(&role).Error
	return role, dbErr
}

func (r *RoleSevc) Update(id uint, role *models.Role) (*models.Role, error) {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	dbErr := db.Model(role).Where("id = ?", id).Update(&role).Find(&role).Error
	return role, dbErr
}

func (r *RoleSevc) Delete(id uint) error {
	db, err := r.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	dbErr := db.Where("id = ?", id).Delete(models.Role{}).Error
	return dbErr
}
