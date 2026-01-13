package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	r := gin.Default()

	r.GET("/health" ,func(c *gin.Context){
		c.JSON(200 , gin.H{
			"status": "ok",
			"message": "task manager api is running",
		})
	} )


	log.Println("server is running on port http://localhost:8080")

	r.Run(":8080")
}
