package repository

import (
	"../config"
	models "../model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repo struct {
	config *config.DatabaseConfigurations
}

func NewMySqlRepository(cf *config.DatabaseConfigurations) *Repo {
	return &Repo{
		config: cf,
	}
}

func (r *Repo) GetConnection() (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", r.config.DBUser, r.config.DBPassword, r.config.DBName)
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

func (r *Repo) AutoMigration() {
	db, err := r.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	user := models.User{}
	role := models.Role{}
	quiz := models.Quiz{}
	lang := models.Language{}
	cate := models.Category{}
	timing := models.Timing{}
	rating := models.Rating{}
	historyDate := models.HistoryDate{}
	history := models.History{}
	question := models.Question{}
	choice := models.Choice{}
	state := models.State{}

	db.AutoMigrate(user, role, quiz, lang, cate, timing, rating, historyDate, history, question, choice, state)
	db.Model(quiz).AddForeignKey("created_by", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(quiz).AddForeignKey("language_id", "languages(id)", "RESTRICT", "RESTRICT")
	db.Model(quiz).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(quiz).AddForeignKey("timing_id", "timings(id)", "RESTRICT", "RESTRICT")
	db.Model(quiz).AddForeignKey("state_id", "states(id)", "RESTRICT", "RESTRICT")

	db.Model(rating).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(rating).AddForeignKey("quiz_id", "quizzes(id)", "RESTRICT", "RESTRICT")
	db.Model(history).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(history).AddForeignKey("quiz_id", "quizzes(id)", "RESTRICT", "RESTRICT")
	db.Model(history).AddForeignKey("history_date_id", "history_dates(id)", "RESTRICT", "RESTRICT")
	db.Model(question).AddForeignKey("quiz_id", "quizzes(id)", "RESTRICT", "RESTRICT")
	db.Model(choice).AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")

	db.Table("user_roles").AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Table("user_roles").AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
}
