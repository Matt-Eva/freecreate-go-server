-- migrate:up
    ALTER TABLE users
        ALTER COLUMN reading_history SET NOT NULL;

-- migrate:down

