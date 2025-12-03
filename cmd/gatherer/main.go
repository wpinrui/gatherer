package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpinrui/gatherer/internal/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/upload", handlers.UploadFile)

	log.Println("Gatherer - Starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
