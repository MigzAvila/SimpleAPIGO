-- Filename new_migrations/000003_add_schools_indexes.down.sql

DROP INDEX IF EXISTS school_name_idx;
DROP INDEX IF EXISTS school_level_idx;
DROP INDEX IF EXISTS school_mode_idx;
