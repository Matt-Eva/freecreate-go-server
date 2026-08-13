#!/bin/bash

export $(grep -v '^#' .env | xargs)

echo "running pg content migration"

echo $PG_CONTENT_DB_ONE_URL

dbmate -d "./internal/db/pg_content/migrations" -s "./internal/db/pg_content/schema.sql" --url $PG_CONTENT_DB_ONE_URL migrate