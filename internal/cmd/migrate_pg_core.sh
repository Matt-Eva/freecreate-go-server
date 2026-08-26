#!/bin/bash
export $(grep -v '^#' .env | xargs)

echo "running pg core migration"

echo $PG_MAIN_DB_URL

dbmate -d "./internal/db/pg_core/migrations" -s "./internal/db/pg_core/schema.sql" --url $PG_MAIN_DB_URL migrate