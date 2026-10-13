-- migrate:up
ALTER TABLE chapters
    ADD COLUMN content_db smallint NOT NULL;

-- migrate:down
ALTER TABLE chapters
    DROP COLUMN content_db;

