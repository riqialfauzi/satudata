-- Migration: Create groups table
-- Description: Tabel untuk kategori/grup dataset

CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug VARCHAR(200) NOT NULL UNIQUE,
    name VARCHAR(300) NOT NULL,
    description TEXT,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_groups_slug ON groups(slug);
CREATE INDEX idx_groups_deleted_at ON groups(deleted_at);

-- Seed data grup/kategori
INSERT INTO groups (slug, name, description) VALUES
    ('statistik', 'Statistik', 'Data statistik sektoral dan nasional'),
    ('geospasial', 'Geospasial', 'Data geospasial dan pemetaan'),
    ('ekonomi', 'Ekonomi', 'Data ekonomi dan keuangan'),
    ('sosial', 'Sosial', 'Data sosial dan kependudukan'),
    ('infrastruktur', 'Infrastruktur', 'Data infrastruktur dan pembangunan');
