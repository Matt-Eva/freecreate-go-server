-- migrate:up
CREATE TABLE chapters (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT uuidv7(),
    user_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    writing_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    chapter_number INT NOT NULL,
    published BOOLEAN NOT NULL DEFAULT false,
    published_before BOOLEAN NOT NULL DEFAULT false,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES creators(id) ON DELETE CASCADE,
    FOREIGN KEY (writing_id) REFERENCES writings(id) ON DELETE CASCADE
);

CREATE INDEX idx_chapters_uuid ON chapters(uuid);
CREATE INDEX idx_chapters_writing_id ON chapters(writing_id);

-- migrate:down
DROP TABLE chapters;