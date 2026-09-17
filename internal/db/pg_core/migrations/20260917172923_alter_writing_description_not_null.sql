-- migrate:up

    ALTER TABLE writings
        ALTER COLUMN description SET NOT NULL;

-- migrate:down

