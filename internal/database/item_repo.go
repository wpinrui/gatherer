package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID           uuid.UUID
	OriginalName string
	StoredName   string
	FilePath     string
	FileSize     int64
	MimeType     sql.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(item *Item) error {
	query := `
		INSERT INTO items (id, original_name, stored_name, file_path, file_size, mime_type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Exec(query,
		item.ID, item.OriginalName, item.StoredName, item.FilePath,
		item.FileSize, item.MimeType, item.CreatedAt, item.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}
	return nil
}

func (r *ItemRepository) GetByID(id uuid.UUID) (*Item, error) {
	query := `
		SELECT id, original_name, stored_name, file_path, file_size, mime_type, created_at, updated_at
		FROM items WHERE id = $1
	`
	item := &Item{}
	err := r.db.QueryRow(query, id).Scan(
		&item.ID, &item.OriginalName, &item.StoredName, &item.FilePath,
		&item.FileSize, &item.MimeType, &item.CreatedAt, &item.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("item not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}
	return item, nil
}

func (r *ItemRepository) List() ([]*Item, error) {
	query := `
		SELECT id, original_name, stored_name, file_path, file_size, mime_type, created_at, updated_at
		FROM items ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list items: %w", err)
	}
	defer rows.Close()

	var items []*Item
	for rows.Next() {
		item := &Item{}
		err := rows.Scan(
			&item.ID, &item.OriginalName, &item.StoredName, &item.FilePath,
			&item.FileSize, &item.MimeType, &item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}
