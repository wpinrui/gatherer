package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpinrui/gatherer/internal/handlers"
	"github.com/wpinrui/gatherer/internal/storage"
)

const uploadDir = "./uploads"

func main() {
	fileStorage, err := storage.NewLocalStorage(uploadDir)
	if err != nil {
		log.Fatal("Failed to initialize storage:", err)
	}

	uploadHandler := handlers.NewUploadHandler(fileStorage)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/upload", uploadHandler.Handle)

	log.Println("Gatherer - Starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
