-- migrate:up

    ALTER TABLE users
        ALTER COLUMN user_handle SET DEFAULT '';

    UPDATE users
        SET user_handle = ''
        WHERE user_handle IS NULL;

    ALTER TABLE users
        ALTER COLUMN user_handle SET NOT NULL;


-- migrate:down

