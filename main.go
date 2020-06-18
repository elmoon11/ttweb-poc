package main

import (
	"net/http"

	"github.com/elmoon11/ttweb-poc/controller"
	"github.com/elmoon11/ttweb-poc/repository"
	"github.com/elmoon11/ttweb-poc/service"
	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
)

var (
	taskRepository repository.TaskRepository = repository.NewTaskRepository()
	taskService    service.TaskService       = service.New(taskRepository)
	taskController controller.TaskController = controller.New(taskService)
)

func main() {
	server := gin.Default()

	defer taskRepository.CloseDB()

	server.Use(gin.Recovery())
	server.Use(gindump.Dump())

	routeGroup := server.Group("/api")
	{
		routeGroup.GET("/tasks", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, taskController.FindAll())
		})

		routeGroup.POST("/tasks", func(ctx *gin.Context) {
			err := taskController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "saved"})
			}
		})

		routeGroup.PUT("/tasks/:id", func(ctx *gin.Context) {
			err := taskController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
			}
		})

		routeGroup.DELETE("/tasks/:id", func(ctx *gin.Context) {
			err := taskController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
			}
		})

	}

	server.Run(":8080")

}
