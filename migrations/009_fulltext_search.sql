-- Migration: Full-text search support
-- Description: Menambahkan tsvector column dan GIN index untuk full-text search

-- Tambah kolom tsvector pada tabel releases
ALTER TABLE releases ADD COLUMN IF NOT EXISTS search_vector tsvector;

-- Buat function untuk update search_vector
CREATE OR REPLACE FUNCTION releases_search_vector_update() RETURNS trigger AS $$
BEGIN
    NEW.search_vector := to_tsvector('indonesian', COALESCE(NEW.title, '') || ' ' || COALESCE(NEW.description, ''));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk auto-update search_vector
DROP TRIGGER IF EXISTS trg_releases_search_vector ON releases;
CREATE TRIGGER trg_releases_search_vector
    BEFORE INSERT OR UPDATE ON releases
    FOR EACH ROW EXECUTE FUNCTION releases_search_vector_update();

-- Update existing rows
UPDATE releases SET search_vector = to_tsvector('indonesian', COALESCE(title, '') || ' ' || COALESCE(description, ''));

-- GIN index untuk full-text search
CREATE INDEX IF NOT EXISTS idx_releases_search_vector ON releases USING GIN(search_vector);
