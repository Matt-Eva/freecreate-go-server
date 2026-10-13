-- migrate:up
CREATE TABLE chapter_content (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    chapter_uuid UUID NOT NULL,
    json_content JSONB NOT NULL DEFAULT '{}'::jsonb,
    html_content TEXT NOT NULL DEFAULT '',
);

-- migrate:down
DROP TABLE chapter_content;