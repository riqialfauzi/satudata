-- Migration: Create releases and related tables
-- Description: Tabel untuk rilis data (dataset dan artikel)

CREATE TABLE releases (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(500) NOT NULL,
    slug VARCHAR(500) NOT NULL UNIQUE,
    description TEXT,
    release_type VARCHAR(50) NOT NULL CHECK (release_type IN ('dataset', 'article', 'infographic')),
    status VARCHAR(50) NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'published', 'archived')),
    year INT NOT NULL,
    cover_image_url TEXT,
    tags TEXT[] DEFAULT '{}',
    view_count INT NOT NULL DEFAULT 0,
    published_at TIMESTAMP WITH TIME ZONE,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_releases_slug ON releases(slug);
CREATE INDEX idx_releases_type ON releases(release_type);
CREATE INDEX idx_releases_status ON releases(status);
CREATE INDEX idx_releases_year ON releases(year);
CREATE INDEX idx_releases_created_by ON releases(created_by);
CREATE INDEX idx_releases_deleted_at ON releases(deleted_at);
CREATE INDEX idx_releases_published_at ON releases(published_at);
