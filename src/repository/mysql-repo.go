package repository

import (
	"../config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	cf = config.NewConfiguration()
)


type repo struct {
	config *config.DatabaseConfigurations
}


func NewMySqlRepository(cf *config.DatabaseConfigurations) *repo  {
	return &repo{
		config: cf,
	}
}

func (r *repo) GetConnection() (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",r.config.DBUser,r.config.DBPassword,r.config.DBName)
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	//defer db.Close()
	//er := db.DB().Ping()
	//if er == nil {
	//	fmt.Println("Hello world")
	//}
	return db, nil
}
