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
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_creators_user_id ON creators(user_id);
CREATE INDEX idx_creators_uuid ON creators(uuid);

CREATE UNIQUE INDEX idx_creators_name_user_id ON creators(user_id, name);

CREATE INDEX idx_creators_name_search ON creators USING GIN(creator_name_search_vector);


-- migrate:down

DROP TABLE creators;
