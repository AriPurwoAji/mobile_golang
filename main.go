package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ari Purwo Aji!",
		})
	})

	r.GET("/api/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   1,
			"nama": "Ari",
			"role": "admin",
		})
	})

	r.Run(":8080")
}
