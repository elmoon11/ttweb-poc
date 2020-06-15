package entities

import (
	"time"
)

type Task struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"required" gorm:"type:varchar(200)"`
	Description string    `json:"description" gorm:"type:varchar(250)"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	StartTime   time.Time `json:"startTime" gorm:"type:time"`
	EndTime     time.Time `json:"endTime" gorm:"type:time"`
	//Tags        []Tag     `json:"tags" gorm:"foreignkey:TagId"`
}
