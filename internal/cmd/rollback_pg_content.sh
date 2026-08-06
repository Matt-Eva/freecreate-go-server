#!/bin/bash
echo "rolling back pg content"

dbmate -d "./db/pg_content/migrations" -s "./db/pg_content/schema.sql" --url "postgres://matte:code@localhost:5432/freecreate_go_writing_content?sslmode=disable" rollback