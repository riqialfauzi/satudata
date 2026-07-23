-- Migration: Create organizations table
-- Description: Tabel untuk organisasi/unit kerja

CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug VARCHAR(200) NOT NULL UNIQUE,
    name VARCHAR(300) NOT NULL,
    description TEXT,
    image_url TEXT,
    website VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_organizations_slug ON organizations(slug);
CREATE INDEX idx_organizations_deleted_at ON organizations(deleted_at);

-- Seed data organisasi
INSERT INTO organizations (slug, name, description) VALUES
    ('kementerian-komunikasi-dan-informatika', 'Kementerian Komunikasi dan Informatika', 'Kementerian Komunikasi dan Informatika Republik Indonesia'),
    ('badan-pusat-statistik', 'Badan Pusat Statistik', 'Badan Pusat Statistik Republik Indonesia'),
    ('kementerian-dalam-negeri', 'Kementerian Dalam Negeri', 'Kementerian Dalam Negeri Republik Indonesia');
