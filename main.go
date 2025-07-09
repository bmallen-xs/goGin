package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	Port = "8080"
)

func main() {
	r := gin.Default()

	r.GET("/", indexEndpoint)

	port := Port

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}

func indexEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "good", "code": 200})
}
