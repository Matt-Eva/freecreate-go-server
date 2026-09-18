-- migrate:up

ALTER TABLE users
ADD CONSTRAINT unique_user_handle UNIQUE (user_handle);

-- migrate:down

