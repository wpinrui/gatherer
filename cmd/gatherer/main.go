package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpinrui/gatherer/internal/database"
	"github.com/wpinrui/gatherer/internal/handlers"
	"github.com/wpinrui/gatherer/internal/storage"
)

const uploadDir = "./uploads"

func main() {
	db, err := database.Connect(database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "gatherer",
		Password: "gatherer",
		DBName:   "gatherer",
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	itemRepo := database.NewPostgresItemRepository(db)
	fileStorage, err := storage.NewLocalStorage(uploadDir, itemRepo)
	if err != nil {
		log.Fatal("Failed to initialize storage:", err)
	}

	uploadHandler := handlers.NewUploadHandler(fileStorage)
	itemsHandler := handlers.NewItemsHandler(itemRepo)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/upload", uploadHandler.Handle)
	r.GET("/items", itemsHandler.List)

	log.Println("Gatherer - Starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
