package main

import (
	"log"

	"task-manager/internal/config"
	"task-manager/internal/db"
	"task-manager/internal/routes"
	"github.com/gin-gonic/gin"

)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDatabase(cfg)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"env":    cfg.AppEnv,
		})
	})

	routes.RegisterAuthRoutes(r)
	routes.RegisterProtectedRoutes(r)
	routes.RegisterTaskRoutes(r)

	log.Printf("Server running on http://localhost:%s", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
