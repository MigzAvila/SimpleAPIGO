-- Filename new_migrations/000003_add_schools_indexes.up.sql

CREATE INDEX IF NOT EXISTS school_name_idx ON schools USING GIN(to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS school_level_idx ON schools USING GIN(to_tsvector('simple', level));
CREATE INDEX IF NOT EXISTS school_mode_idx ON schools USING GIN(mode);