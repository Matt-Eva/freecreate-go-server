-- migrate:up
CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT uuidv7(),
    email TEXT UNIQUE NOT NULL CHECK (length(email) < 255),
    username TEXT CHECK (length(username) < 100),
    user_handle TEXT CHECK (length(user_handle) < 100),
    reading_history BOOLEAN DEFAULT false,
    is_adult BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_uuid ON users(uuid);

-- migrate:down

DROP TABLE users;

