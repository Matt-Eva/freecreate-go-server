-- migrate:up
ALTER TABLE chapters
    ADD COLUMN content_db TEXT NOT NULL;

-- migrate:down
ALTER TABLE chapters
    DROP COLUMN content_db;

