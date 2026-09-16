package config

import (
	"context"
	"fmt"
	"freecreate/internal/lib/logger"
	"net/url"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgContentPools struct {
	PgContentOne *pgxpool.Pool
}

func ConfigPgx(ctx context.Context, environment string) (*pgxpool.Pool, PgContentPools, error) {
	corePool, coreErr := connectPgx(ctx, environment, os.Getenv("PG_MAIN_DB_URL"), "./db/pg_core/migrations")
	if coreErr != nil {
		logger.Log(coreErr)
		return nil, PgContentPools{}, coreErr
	}

	contentOnePool, contentErr := connectPgx(ctx, environment, os.Getenv("PG_CONTENT_DB_ONE_URL"), "./db/pg_content/migrations")
	if contentErr != nil {
		logger.Log(contentErr)
		return nil, PgContentPools{}, contentErr
	}

	pgContentPools := PgContentPools{
		PgContentOne: contentOnePool,
	}

	return corePool, pgContentPools, nil
}

func connectPgx(ctx context.Context, environment string, connEnv string, migrationsDir string) (*pgxpool.Pool, error) {

	pgxPool, err := pgxpool.New(ctx, connEnv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	msg := fmt.Sprintf("successful connection to %s!", migrationsDir)
	fmt.Println(msg)

	if environment == "PRODUCTION" {
		migrationErr := runDbmateMigrations(connEnv, environment, migrationsDir)
		if migrationErr != nil {
			logger.Log(migrationErr)
			return nil, migrationErr
		}
	} else {
		msg := fmt.Sprintf("migrations not run, environment is %s, not PRODUCTION", environment)
		fmt.Println(msg)
	}

	return pgxPool, nil
}

func runDbmateMigrations(connString string, environment string, migrationsDir string) error {
	u, _ := url.Parse(connString)
	db := dbmate.New(u)

	db.MigrationsDir = []string{migrationsDir}

	if environment == "PRODUCTION" {
		db.AutoDumpSchema = false
	}

	migrationErr := db.Migrate()
	if migrationErr != nil {
		fmt.Fprintf(os.Stderr, "unable to run dbmate migration: %v\n", migrationErr)
		return migrationErr
	}

	return nil
}
