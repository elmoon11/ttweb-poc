package model

import (
	_ "time"
)

type Task struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Created     string `json:"created"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}
