package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := SetupRouter()
	if err := r.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
