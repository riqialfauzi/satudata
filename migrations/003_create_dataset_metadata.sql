-- Migration: Create dataset_metadata table
-- Description: Metadata spesifik untuk rilis tipe dataset

CREATE TABLE dataset_metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    release_id UUID NOT NULL REFERENCES releases(id) ON DELETE CASCADE,
    file_url TEXT NOT NULL,
    file_format VARCHAR(50) NOT NULL CHECK (file_format IN ('csv', 'json', 'xlsx', 'parquet')),
    file_size BIGINT NOT NULL DEFAULT 0,
    row_count INT,
    column_count INT,
    columns JSONB,
    data_source VARCHAR(500),
    data_period_start DATE,
    data_period_end DATE,
    update_frequency VARCHAR(100),
    is_geospatial BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_dataset_metadata_release_id ON dataset_metadata(release_id);
CREATE INDEX idx_dataset_metadata_format ON dataset_metadata(file_format);
