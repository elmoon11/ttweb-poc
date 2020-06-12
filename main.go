package main

import (
	"net/http"

	"github.com/elmoon11/ttweb-poc/controller"
	"github.com/elmoon11/ttweb-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	taskService    service.TaskService       = service.New()
	taskController controller.TaskController = controller.New(taskService)
)

func main() {
	server := gin.Default()

	server.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, taskController.FindAll())
	})

	server.POST("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, taskController.Save(ctx))
	})

	server.Run(":8080")

}
