-- Migration: Create article_metadata table
-- Description: Metadata spesifik untuk rilis tipe artikel

CREATE TABLE article_metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    release_id UUID NOT NULL REFERENCES releases(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    excerpt TEXT,
    author_name VARCHAR(255),
    reading_time_minutes INT,
    is_featured BOOLEAN NOT NULL DEFAULT false,
    category VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_article_metadata_release_id ON article_metadata(release_id);
CREATE INDEX idx_article_metadata_category ON article_metadata(category);
CREATE INDEX idx_article_metadata_featured ON article_metadata(is_featured);
