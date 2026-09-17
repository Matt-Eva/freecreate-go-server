-- migrate:up


    ALTER TABLE users
        ALTER COLUMN username SET DEFAULT '';

    UPDATE users
        SET username = ''
        WHERE username IS NULL;

    ALTER TABLE users
        ALTER COLUMN username SET NOT NULL;


-- migrate:down

