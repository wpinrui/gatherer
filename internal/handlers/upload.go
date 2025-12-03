package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	UploadDir     = "./uploads"
	MaxUploadSize = 50 << 20 // 50 MB
)

type UploadResponse struct {
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	Size      int64  `json:"size"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
}

func init() {
	if err := os.MkdirAll(UploadDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create upload directory: %v", err))
	}
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no file provided",
		})
		return
	}

	if file.Size > MaxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("file too large, max size is %d MB", MaxUploadSize>>20),
		})
		return
	}

	id := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	storedName := id + ext
	destPath := filepath.Join(UploadDir, storedName)

	if err := c.SaveUploadedFile(file, destPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save file",
		})
		return
	}

	response := UploadResponse{
		ID:        id,
		Filename:  file.Filename,
		Size:      file.Size,
		Path:      destPath,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	c.JSON(http.StatusCreated, response)
}
