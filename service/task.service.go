package service

import (
	"github.com/elmoon11/ttweb-poc/entities"
	"github.com/elmoon11/ttweb-poc/repository"
)

type TaskService interface {
	Update(entities.Task)
	Delete(entities.Task)
	Save(entities.Task) entities.Task
	FindAll() []entities.Task
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func New(repository repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: repository,
	}
}

func (service *taskService) Update(task entities.Task) {
	service.taskRepository.Update(task)
}

func (service *taskService) Delete(task entities.Task) {
	service.taskRepository.Delete(task)
}

func (service *taskService) Save(task entities.Task) entities.Task {
	service.taskRepository.Save(task)
	return task
}

func (service *taskService) FindAll() []entities.Task {
	return service.taskRepository.FindAll()
}
