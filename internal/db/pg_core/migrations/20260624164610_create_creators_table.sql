-- migrate:up
CREATE TABLE creators (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT uuidv7(),
    user_id BIGINT NOT NULL,
    creator_language REGCONFIG NOT NULL DEFAULT 'english',
    name TEXT NOT NULL CHECK (length(name) < 100),
    creator_handle TEXT UNIQUE CHECK (length(creator_handle) < 100),
    creator_name_search_vector tsvector GENERATED ALWAYS AS (
        to_tsvector(creator_language, name) || to_tsvector(creator_language, creator_handle)
    ) STORED,
    topics TEXT ARRAY NOT NULL DEFAULT ARRAY[]::TEXT[],
    tags TEXT ARRAY NOT NULL DEFAULT ARRAY[]::TEXT[] CHECK (cardinality(tags) <= 20), 
    writing_types TEXT ARRAY NOT NULL DEFAULT ARRAY[]::TEXT[],
    rank BIGINT NOT NULL DEFAULT 0,
    rel_rank BIGINT NOT NULL DEFAULT 0,
    donations BIGINT NOT NULL DEFAULT 0,
    supporters BIGINT NOT NULL DEFAULT 0,
    followers BIGINT NOT NULL DEFAULT 0,
    subscribers BIGINT NOT NULL DEFAULT 0,
    views BIGINT NOT NULL DEFAULT 0,
    flags BIGINT NOT NULL DEFAULT 0,
    rank_tracker BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_published TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_creators_user_id ON creators(user_id);
CREATE INDEX idx_creators_uuid ON creators(uuid);
CREATE INDEX idx_creators_rank ON creators(rank);
CREATE INDEX idx_creators_rel_rank ON creators(rel_rank);
CREATE INDEX idx_creators_last_published ON creators(last_published);

CREATE UNIQUE INDEX idx_creators_name_user_id ON creators(user_id, name);

CREATE INDEX idx_creator_topics ON creators USING GIN(topics);
CREATE INDEX idx_creator_tags ON creators USING GIN(tags);
CREATE INDEX idx_creator_writing_types ON creators USING GIN(writing_types);
CREATE INDEX idx_creators_name_search ON creators USING GIN(creator_name_search_vector);


-- migrate:down

DROP TABLE creators;
