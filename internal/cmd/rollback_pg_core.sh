#!/bin/bash
export $(grep -v '^#' .env | xargs)

echo "rolling back pg core"

echo $PG_MAIN_DB_URL 

dbmate -d "./internal/db/pg_core/migrations" -s "./internal/db/pg_core/schema.sql" --url $PG_MAIN_DB_URL rollback