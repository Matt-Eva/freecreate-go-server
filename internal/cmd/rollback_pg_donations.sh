#!/bin/bash

export $(grep -v '^#' .env | xargs)

echo "running pg content migration"

echo $PG_DONATIONS_DB_ONE_URL

dbmate -d "./internal/db/pg_content/migrations" -s "./internal/db/pg_content/schema.sql" --url $PG_DONATIONS_DB_ONE_URL rollback