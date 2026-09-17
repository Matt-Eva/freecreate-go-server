-- migrate:up

    ALTER TABLE users
        ALTER COLUMN is_adult SET NOT NULL;

-- migrate:down

