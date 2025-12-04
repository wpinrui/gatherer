-- Gatherer database schema

CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY,
    original_name VARCHAR(255) NOT NULL,
    stored_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(512) NOT NULL,
    file_size BIGINT NOT NULL,
    mime_type VARCHAR(127),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index for listing items by creation date
CREATE INDEX IF NOT EXISTS idx_items_created_at ON items(created_at DESC);
