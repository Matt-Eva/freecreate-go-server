package query_handlers

import (
	"context"
	"freecreate/internal/cache"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func HandleSearch(ctx context.Context, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, searchType string, writingType string, datePosted string, topics []string, tags []string)*api_error.Error{

	cache.GetSearch(ctx, valkeyClient, searchType, writingType, datePosted, topics, tags)

	return nil
}