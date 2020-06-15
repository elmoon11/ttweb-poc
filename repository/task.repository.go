package repository

import (
	"github.com/elmoon11/ttweb-poc/entities"
	"github.com/jinzhu/gorm"
)

type TaskRepository interface {
	Save(task entities.Task)
	Update(task entities.Task)
	Delete(task entities.Task)
	FindAll() []entities.Task
}

type database struct {
	conn *gorm.DB
}

func New() TaskRepository {
	db, err := gorm.Open("sqlite3", "temp.db")
	if err != nil {
		panic("Db connection failed")
	}
	//db.AutoMigrate(&entities.Task)
	return &database{
		conn: db,
	}
}
