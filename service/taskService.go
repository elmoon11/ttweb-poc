package service

import (
	"github.com/elmoon11/ttweb-poc/model"
)

type TaskService interface {
	Save(model.Task) model.Task
	FindAll() []model.Task
}

type taskService struct {
	tasks []model.Task
}

func New() TaskService {
	return &taskService{}
}

func (service *taskService) Save(task model.Task) model.Task {
	service.tasks = append(service.tasks, task)
	return task
}

func (service *taskService) FindAll() []model.Task {
	return service.tasks
}
