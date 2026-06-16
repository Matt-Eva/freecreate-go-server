package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)



func ConfigPgxMainDB(ctx context.Context) *pgxpool.Pool{
		user, pwd, host, port, db, ssl :=  os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST") , os.Getenv("PG_PORT"), os.Getenv("PG_MAIN_DB"), os.Getenv("PG_SSL")
		connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pwd, host, port, db, ssl)
	
		pgxMainPool, err := pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("successful connection to pgx Main db!")
		return pgxMainPool
}

func ConfigPgxContentDB(ctx context.Context){

}