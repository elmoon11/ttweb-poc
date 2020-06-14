package controller

import (
	"github.com/elmoon11/ttweb-poc/model"
	"github.com/elmoon11/ttweb-poc/service"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	FindAll() []model.Task
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

func (c *controller) FindAll() []model.Task {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var task model.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		return err
	}
	c.service.Save(task)
	return nil
}
