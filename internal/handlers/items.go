package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpinrui/gatherer/internal/database"
)

// ItemsHandler handles item listing requests.
type ItemsHandler struct {
	repo database.ItemRepository
}

// NewItemsHandler creates a new ItemsHandler.
func NewItemsHandler(repo database.ItemRepository) *ItemsHandler {
	return &ItemsHandler{repo: repo}
}

// ItemResponse is the API response for a single item.
type ItemResponse struct {
	ID           string `json:"id"`
	OriginalName string `json:"original_name"`
	FileSize     int64  `json:"file_size"`
	MimeType     string `json:"mime_type,omitempty"`
	CreatedAt    string `json:"created_at"`
}

// List returns all items.
func (h *ItemsHandler) List(c *gin.Context) {
	items, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve items",
		})
		return
	}

	response := make([]ItemResponse, 0, len(items))
	for _, item := range items {
		mimeType := ""
		if item.MimeType.Valid {
			mimeType = item.MimeType.String
		}
		response = append(response, ItemResponse{
			ID:           item.ID.String(),
			OriginalName: item.OriginalName,
			FileSize:     item.FileSize,
			MimeType:     mimeType,
			CreatedAt:    item.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	c.JSON(http.StatusOK, response)
}
