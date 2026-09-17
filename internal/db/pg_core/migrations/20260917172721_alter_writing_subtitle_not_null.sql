-- migrate:up


    ALTER TABLE writings
        ALTER COLUMN subtitle SET DEFAULT '';

    UPDATE writings
        SET subtitle = ''
        WHERE subtitle IS NULL;

    ALTER TABLE writings
        ALTER COLUMN subtitle SET NOT NULL;

-- migrate:down

