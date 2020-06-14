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
		err := taskController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "saved"})
		}
	})

	server.Run(":8080")

}
