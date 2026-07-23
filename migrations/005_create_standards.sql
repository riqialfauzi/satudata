-- Migration: Create standards table
-- Description: Tabel untuk menyimpan standar data tahunan

CREATE TABLE standards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(500) NOT NULL,
    description TEXT,
    year INT NOT NULL,
    file_url TEXT,
    file_size BIGINT DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'archived', 'draft')),
    version VARCHAR(50) NOT NULL DEFAULT '1.0',
    is_current BOOLEAN NOT NULL DEFAULT false,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(year, version)
);

CREATE INDEX idx_standards_year ON standards(year);
CREATE INDEX idx_standards_status ON standards(status);
CREATE INDEX idx_standards_is_current ON standards(is_current);
CREATE INDEX idx_standards_deleted_at ON standards(deleted_at);
