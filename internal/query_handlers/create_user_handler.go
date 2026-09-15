package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateUserParams struct {
	Email string
}

type CreatedUser struct {
	UserId int64
}

func HandleCreateUser(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool, userParams CreateUserParams)(CreatedUser, *api_error.Error) {
	var createdUser CreatedUser

	userId, createUserErr := pg_core_queries.CreateUser(ctx, pgCoreQueries, pgCore, userParams.Email)
	if createUserErr != nil {
		return createdUser, createUserErr
	}

	createdUser.UserId = userId
	
	return createdUser, nil
}
