-- migrate:up
ALTER TABLE writings
    ADD COLUMN is_adult BOOLEAN DEFAULT true;

-- migrate:down
ALTER TABLE writings
    DROP COLUMN is_adult;
