package model

import (
	_ "time"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}
