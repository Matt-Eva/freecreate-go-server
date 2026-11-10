-- migrate:up
ALTER TABLE chapter_content
    ADD COLUMN user_uuid UUID NOT NULL;

-- migrate:down
ALTER TABLE chapter_content
    DROP COLUMN user_uuid;

