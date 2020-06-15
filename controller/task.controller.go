package controller

import (
	"github.com/elmoon11/ttweb-poc/entities"
	"github.com/elmoon11/ttweb-poc/service"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	FindAll() []entities.Task
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.TaskService
}

func New(service service.TaskService) TaskController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entities.Task {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var task entities.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		return err
	}
	c.service.Save(task)
	return nil
}
