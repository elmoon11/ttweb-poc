package service

import (
	"github.com/elmoon11/ttweb-poc/entities"
)

type TaskService interface {
	Save(entities.Task) entities.Task
	FindAll() []entities.Task
}

type taskService struct {
	tasks []entities.Task
}

func New() TaskService {
	return &taskService{}
}

func (service *taskService) Save(task entities.Task) entities.Task {
	service.tasks = append(service.tasks, task)
	return task
}

func (service *taskService) FindAll() []entities.Task {
	return service.tasks
}
