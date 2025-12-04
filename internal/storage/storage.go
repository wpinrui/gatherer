package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
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

// LocalStorage implements FileStorage using the local filesystem.
type LocalStorage struct {
	baseDir string
	files   map[string]*FileMetadata
}

// NewLocalStorage creates a new LocalStorage instance.
func NewLocalStorage(baseDir string) (*LocalStorage, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &LocalStorage{
		baseDir: baseDir,
		files:   make(map[string]*FileMetadata),
	}, nil
}

// Save stores a file and returns its metadata.
func (s *LocalStorage) Save(filename string, content io.Reader, size int64) (*FileMetadata, error) {
	id := uuid.New().String()
	ext := filepath.Ext(filename)
	storedName := id + ext
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

	metadata := &FileMetadata{
		ID:           id,
		OriginalName: filename,
		StoredName:   storedName,
		Size:         written,
		Path:         destPath,
		CreatedAt:    time.Now().UTC(),
	}

	s.files[id] = metadata

	return metadata, nil
}

// Get retrieves file metadata by ID.
func (s *LocalStorage) Get(id string) (*FileMetadata, error) {
	metadata, exists := s.files[id]
	if !exists {
		return nil, fmt.Errorf("file not found: %s", id)
	}
	return metadata, nil
}

// Delete removes a file by ID.
func (s *LocalStorage) Delete(id string) error {
	metadata, exists := s.files[id]
	if !exists {
		return fmt.Errorf("file not found: %s", id)
	}

	if err := os.Remove(metadata.Path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	delete(s.files, id)
	return nil
}
