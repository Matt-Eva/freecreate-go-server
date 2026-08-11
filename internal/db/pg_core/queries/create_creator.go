package pg_core_queries

import (
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NewCreator struct{

}

func CreateCreator(pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries)(newCreator, *api_error.Error){
	query := pgCoreQueries.CreateCreator()
}