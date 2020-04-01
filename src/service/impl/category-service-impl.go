package impl

import (
	models "../../model"
	"../../repository"
	"../../service"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Sevc struct {
	db *repository.Repo
}



func NewCategoryService(db *repository.Repo) service.CategoryService {
	return &Sevc{
		db: db,
	}
}

func (s *Sevc) FindAll() []models.Category {
	fmt.Println("Category FindALl service")
	return nil
}

func (s *Sevc) Save(category models.Category) models.Category {
	db, err := s.db.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.DB().Close()
	ctx := context.Background()
	tx, err1 := db.DB().BeginTx(ctx,&sql.TxOptions{Isolation: sql.LevelDefault})
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	sql := "INSERT INTO categories(name,created_at) VALUES (?, ?)"
	res, execErr := tx.ExecContext(ctx,sql,category.Name,time.Now())
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatalf("%v\n%v",execErr,rollBackErr)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println(res.LastInsertId())
	var cate models.Category
	//db.Where("id = ?", res.LastInsertId()).Find(&cate)
	return cate
}


