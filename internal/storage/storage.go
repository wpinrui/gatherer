package storage

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/wpinrui/gatherer/internal/database"
)

// FileMetadata contains information about a stored file.
type FileMetadata struct {
	ID           string    `json:"id"`
	OriginalName string    `json:"original_name"`
	StoredName   string    `json:"stored_name"`
	Size         int64     `json:"size"`
	Path         string    `json:"path"`
	CreatedAt    time.Time `json:"created_at"`
}

// FileStorage defines the interface for file storage operations.
type FileStorage interface {
	Save(filename string, content io.Reader, size int64) (*FileMetadata, error)
	Get(id string) (*FileMetadata, error)
	Delete(id string) error
}

// LocalStorage implements FileStorage using filesystem + database.
type LocalStorage struct {
	baseDir  string
	itemRepo *database.ItemRepository
}

// NewLocalStorage creates a new LocalStorage instance.
func NewLocalStorage(baseDir string, itemRepo *database.ItemRepository) (*LocalStorage, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &LocalStorage{
		baseDir:  baseDir,
		itemRepo: itemRepo,
	}, nil
}

// Save stores a file and returns its metadata.
func (s *LocalStorage) Save(filename string, content io.Reader, size int64) (*FileMetadata, error) {
	id := uuid.New()
	ext := filepath.Ext(filename)
	storedName := id.String() + ext
	destPath := filepath.Join(s.baseDir, storedName)

	destFile, err := os.Create(destPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer destFile.Close()

	written, err := io.Copy(destFile, content)
	if err != nil {
		os.Remove(destPath)
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	now := time.Now().UTC()
	item := &database.Item{
		ID:           id,
		OriginalName: filename,
		StoredName:   storedName,
		FilePath:     destPath,
		FileSize:     written,
		MimeType:     sql.NullString{},
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.itemRepo.Create(item); err != nil {
		os.Remove(destPath)
		return nil, fmt.Errorf("failed to save metadata: %w", err)
	}

	return &FileMetadata{
		ID:           id.String(),
		OriginalName: filename,
		StoredName:   storedName,
		Size:         written,
		Path:         destPath,
		CreatedAt:    now,
	}, nil
}

// Get retrieves file metadata by ID.
func (s *LocalStorage) Get(id string) (*FileMetadata, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	item, err := s.itemRepo.GetByID(uid)
	if err != nil {
		return nil, err
	}

	return &FileMetadata{
		ID:           item.ID.String(),
		OriginalName: item.OriginalName,
		StoredName:   item.StoredName,
		Size:         item.FileSize,
		Path:         item.FilePath,
		CreatedAt:    item.CreatedAt,
	}, nil
}

// Delete removes a file by ID.
func (s *LocalStorage) Delete(id string) error {
	metadata, err := s.Get(id)
	if err != nil {
		return err
	}

	if err := os.Remove(metadata.Path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
