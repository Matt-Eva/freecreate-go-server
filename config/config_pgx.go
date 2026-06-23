package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

func ConfigPgxMainDb(ctx context.Context) *pgxpool.Pool {

	mainDBConnURL := os.Getenv("PG_MAIN_DB_URL")

	pgxMainPool, err := pgxpool.New(ctx, mainDBConnURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("successful connection to pgx Main db!")
	return pgxMainPool
}

func ConfigPgxContentDbOne(ctx context.Context) *pgxpool.Pool {
	contentDbOneConnUrl := os.Getenv("PG_CONTENT_DB_ONE_URL")

	pgxContentDbOnePool, err := pgxpool.New(ctx, contentDbOneConnUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("successful connectino to pgx Content db one!")
	return pgxContentDbOnePool

}

// ============== Just create more content DBs in parallel to create "sharding" for content =======
// func ConfigPgxContentDBTwo(ctx context.Context){

// }
