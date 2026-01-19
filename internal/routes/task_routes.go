package routes

import (
	"task-manager/internal/handlers"
	"task-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine) {
	taskHandler := handlers.NewTaskHandler()

	tasks := r.Group("api/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.POST("", taskHandler.Create)
		tasks.GET("", taskHandler.GetAll)
		tasks.PUT("/:id", taskHandler.Update)
		tasks.DELETE("/:id", taskHandler.Delete)
	}
}
