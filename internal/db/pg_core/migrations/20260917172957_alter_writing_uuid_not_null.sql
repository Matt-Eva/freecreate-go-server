-- migrate:up

    ALTER TABLE writings
        ALTER COLUMN uuid SET NOT NULL;

-- migrate:down

