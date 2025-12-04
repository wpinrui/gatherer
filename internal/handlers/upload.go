package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpinrui/gatherer/internal/storage"
)

const MaxUploadSize = 50 << 20 // 50 MB

// UploadHandler handles file upload requests.
type UploadHandler struct {
	storage storage.FileStorage
}

// NewUploadHandler creates a new UploadHandler with the given storage.
func NewUploadHandler(s storage.FileStorage) *UploadHandler {
	return &UploadHandler{storage: s}
}

// UploadResponse is the API response for file uploads.
type UploadResponse struct {
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	Size      int64  `json:"size"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
}

// Handle processes file upload requests.
func (h *UploadHandler) Handle(c *gin.Context) {
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

	src, err := file.Open()
	if err != nil {
		log.Printf("failed to open uploaded file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read file",
		})
		return
	}
	defer src.Close()

	metadata, err := h.storage.Save(file.Filename, src, file.Size)
	if err != nil {
		log.Printf("failed to save file %s: %v", file.Filename, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save file",
		})
		return
	}

	response := UploadResponse{
		ID:        metadata.ID,
		Filename:  metadata.OriginalName,
		Size:      metadata.Size,
		Path:      metadata.Path,
		CreatedAt: metadata.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	c.JSON(http.StatusCreated, response)
}
