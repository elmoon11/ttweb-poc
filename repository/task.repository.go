package repository

import (
	"github.com/elmoon11/ttweb-poc/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type TaskRepository interface {
	Save(task entities.Task)
	Update(task entities.Task)
	Delete(task entities.Task)
	FindAll() []entities.Task
	CloseDB()
}

type database struct {
	conn *gorm.DB
}

func NewTaskRepository() TaskRepository {
	db, err := gorm.Open("sqlite3", "temp.db")
	if err != nil {
		panic("Db connection failed")
	}
	db.AutoMigrate(&entities.Task{})
	return &database{
		conn: db,
	}
}

func (db *database) CloseDB() {
	err := db.conn.Close()
	if err != nil {
		panic("Failed to close datbase")
	}
}

func (db *database) Save(task entities.Task) {
	db.conn.Create(&task)
}

func (db *database) Update(task entities.Task) {
	db.conn.Update(&task)
}

func (db *database) Delete(task entities.Task) {
	db.conn.Delete(&task)
}

func (db *database) FindAll() []entities.Task {
	var tasks []entities.Task
	db.conn.Set("gorm:auto_preload", true).Find(&tasks)
	return tasks
}
