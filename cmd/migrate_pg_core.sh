#!/bin/bash
echo "running pg core migration"

dbmate -d "./db/pg_core/migrations" -s "./db/pg_core/schema.sql" --url "postgres://matte:code@localhost:5432/freecreate_go?sslmode=disable" migrate