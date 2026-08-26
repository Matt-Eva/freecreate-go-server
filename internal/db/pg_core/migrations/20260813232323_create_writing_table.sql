-- migrate:up
CREATE TABLE writings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    uuid UUID DEFAULT uuidv7(),
    writing_language REGCONFIG NOT NULL DEFAULT 'english',
    title TEXT NOT NULL CHECK (length(title) < 100),
    subtitle TEXT CHECK (length(subtitle) < 100),
    title_search_vector TSVECTOR GENERATED ALWAYS AS (
        to_tsvector(writing_language, title) || to_tsvector(writing_language, subtitle)
    ) STORED,
    description TEXT CHECK (length(description) < 300),
    writing_type TEXT NOT NULL,
    topics TEXT ARRAY NOT NULL DEFAULT ARRAY[]::TEXT[] CHECK (cardinality(topics) <= 3) ,
    tags TEXT ARRAY NOT NULL DEFAULT ARRAY[]::TEXT[] CHECK (cardinality(tags) <= 20), 
    rank BIGINT NOT NULL DEFAULT 0,
    rel_rank BIGINT NOT NULL DEFAULT 0,
    views BIGINT NOT NULL DEFAULT 0,
    list_adds BIGINT NOT NULL DEFAULT 0,
    likes BIGINT NOT NULL DEFAULT 0,
    lib_adds BIGINT NOT NULL DEFAULT 0,
    donations BIGINT NOT NULL DEFAULT 0,
    flags BIGINT NOT NULL DEFAULT 0,
    rank_tracker BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    published BOOLEAN NOT NULL DEFAULT false ,
    published_before BOOLEAN NOT NULL DEFAULT false,
    last_published TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES creators(id) ON DELETE CASCADE
);

CREATE INDEX idx_writings_uuid ON writings(uuid);
CREATE INDEX idx_writings_user_id ON writings(user_id);
CREATE INDEX idx_writings_creator_id ON writings(creator_id);
CREATE INDEX idx_writings_rank ON writings(rank);
CREATE INDEX idx_writings_rel_rank ON writings(rel_rank);
CREATE INDEX idx_writings_last_published ON writings(last_published);

CREATE INDEX idx_writings_topics ON writings USING GIN(topics);
CREATE INDEX idx_writings_tags ON writings USING GIN(tags);
CREATE INDEX idx_writings_title_search ON writings USING GIN(title_search_vector);

-- migrate:down
DROP TABLE writings;
