package routes

import (
	"task-manager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine){
	AuthHandler := handlers.NewAuthHandler()

	auth := r.Group("/auth")
	{
		auth.POST("/register", AuthHandler.Register)
		auth.POST("/login" , AuthHandler.Login)
	}


}