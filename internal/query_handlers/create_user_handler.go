package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleCreateUser(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool, userParams pg_core_types.CreateUserParams) (pg_core_types.CreatedUser, *api_error.Error) {

	user, createUserErr := pg_core_queries.CreateUser(ctx, pgCoreQueries, pgCore, userParams.Email)
	if createUserErr != nil {
		return user, createUserErr
	}

	return user, nil
}
