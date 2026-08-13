-- migrate:up
CREATE TABLE writings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    uuid UUID DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    subtitle VARCHAR(200),
    genres TEXT ARRAY,
    tags TEXT ARRAY, 
    rank BIGINT DEFAULT 0,
    rel_rank BIGINT DEFAULT 0,
    description TEXT,
    views BIGINT DEFAULT 0,
    list_adds BIGINT DEFAULT 0,
    likes BIGINT DEFAULT 0,
    lib_adds BIGINT DEFAULT 0,
    donations BIGINT DEFAULT 0,
    flags BIGINT DEFAULT 0,
    created_at TIMESTAMP,
    published BOOLEAN,
    published_before BOOLEAN,
    last_published TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES creators(id) ON DELETE CASCADE
);

CREATE INDEX idx_writings_uuid ON writings(uuid);
CREATE INDEX idx_writings_user_id ON writings(user_id);
CREATE INDEX idx_writings_creator_id ON writings(creator_id);
CREATE INDEX idx_writings_rank ON writings(rank);
CREATE INDEX idx_writings_rel_rank ON writings(rel_rank);
CREATE INDEX idx_writings_created_at ON writings(created_at);
CREATE INDEX idx_writings_last_published ON writings(last_published);

-- migrate:down
DROP TABLE writings;
