package pg_core_queries

import (
	"fmt"
	"freecreate/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByEmail(pgCoreQueries config.PgCoreQueries, pgxCore *pgxpool.Pool, email string)(int, error){
	query := pgCoreQueries.GetUserByEmail()
	fmt.Println(query)
	return 0, nil
}