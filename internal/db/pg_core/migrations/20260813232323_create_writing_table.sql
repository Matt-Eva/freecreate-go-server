-- migrate:up
CREATE TABLE writings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    uuid DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    genres ARRAY,
    tags ARRAY, 
    rank BIGINT default 0,
    rel_rank BIGINT default 0,
    description TEXT,
    views BIGINT default 0,
    list_adds BIGINT default 0,
    likes BIGINT default 0,
    lib_adds BIGINT default 0,
    donations BIGINT default 0,
    flags BIGINT default 0,
    FOREIGN KEY (user_id) REFERENCES users.id ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES creators.id ON DELETE CASCADE
);

CREATE INDEX idx_writings_uuid ON writings.uuid,
CREATE INDEX idx_writings_user_id ON writings.user_id,
CREATE INDEX idx_writings_creator_id ON writings.creator_id

-- migrate:down
DROP TABLE writings;
