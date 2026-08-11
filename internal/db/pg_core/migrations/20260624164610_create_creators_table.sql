-- migrate:up
CREATE TABLE creators (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid(),
    user_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_creators_user_id ON creators(user_id);
CREATE INDEX idx_creators_uuid ON creators(uuid);

-- migrate:down

DROP TABLE creators;
