package controller

import (
	"github.com/elmoon11/ttweb-poc/entities"
	"github.com/elmoon11/ttweb-poc/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TaskController interface {
	FindAll() []entities.Task
	Save(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	Update(ctx *gin.Context) error
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

func (c *controller) Update(ctx *gin.Context) error {
	var task entities.Task
	err := ctx.Copy().ShouldBindJSON(&task)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	task.ID = id

	c.service.Update(task)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var task entities.Task

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	task.ID = id
	c.service.Delete(task)
	return nil
}
