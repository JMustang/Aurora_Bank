package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "Aurora Bank",
		})
	})

	router.Run(":8080")
}
