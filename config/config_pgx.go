package config

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConfigPgxCoreDb(ctx context.Context) *pgxpool.Pool {

	mainDBConnURL := os.Getenv("PG_MAIN_DB_URL")

	pgxMainPool, err := pgxpool.New(ctx, mainDBConnURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("successful connection to pgx Main db!")

	runPgxCoreMigrations(mainDBConnURL)

	return pgxMainPool
}

func runPgxCoreMigrations(connString string) error {
	u, _:= url.Parse(connString)
	db := dbmate.New(u)

	db.MigrationsDir = []string{"./db/pg_core/migrations"}
	db.SchemaFile = "./db/pg_core"

	migrationErr := db.Migrate()
	if migrationErr != nil {
		fmt.Fprintf(os.Stderr, "unable to run dbmate migration: %v\n", migrationErr)
		os.Exit(1)
	}
	return nil
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
