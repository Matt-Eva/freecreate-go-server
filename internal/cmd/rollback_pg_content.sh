#!/bin/bash

export $(grep -v '^#' .env | xargs)

echo "rolling back pg content"

echo $PG_CONTENT_DB_ONE_URL

dbmate -d "./internal/db/pg_content/migrations" -s "./internal/db/pg_content/schema.sql" --url $PG_CONTENT_DB_ONE_URL rollback