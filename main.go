package main

import (
	"github.com/gin-gonic/gin"
	"github.com/JMustang/Aurora_Bank/internal/database"
)

func main() {

	database.Connect()
	router := gin.Default()

	router.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "Aurora Bank",
		})
	})

	router.Run(":8080")
}
