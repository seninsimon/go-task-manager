package main

import (
	"fmt"
	"log"
	"task-manager/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	cfg := config.LoadConfig()

	r := gin.Default()

	r.GET("/health" ,func(c *gin.Context){
		c.JSON(200 , gin.H{
			"status": "ok",
			"env": cfg.AppEnv,
		})
	} )


	log.Println("server is running on port http://localhost:8080")

	r.Run(":8080")
}
