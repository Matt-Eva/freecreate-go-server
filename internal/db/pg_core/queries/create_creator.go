package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, createCreatorParams pg_core_types.CreateCreatorParams) (pg_core_types.CreatedCreator, *api_error.Error) {
	var createdCreator pg_core_types.CreatedCreator

	query := pgCoreQueries.CreateCreator()

	namedArgs := pgx.NamedArgs{
		"name":           createCreatorParams.Name,
		"creator_handle": createCreatorParams.CreatorHandle,
		"user_id":        createCreatorParams.UserId,
	}

	validateCreatorErr := pg_core_validators.ValidateCreator(namedArgs)
	if validateCreatorErr != nil {
		return createdCreator, validateCreatorErr
	}

	var name string
	var creator_handle string
	var uuid uuid.UUID

	createCreatorErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&name, &uuid, &creator_handle)

	var pgErr *pgconn.PgError
	if errors.As(createCreatorErr, &pgErr) && pgErr.Code == "23505" {
		apiErr := &api_error.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "That creator handle is already in use.",
			Error:   createCreatorErr,
		}

		return createdCreator, apiErr
	} else if createCreatorErr != nil {
		logger.Log(createCreatorErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   createCreatorErr,
		}

		return createdCreator, apiErr
	}

	createdCreator.Name = name
	createdCreator.UUID = uuid
	createdCreator.CreatorHandle = creator_handle

	return createdCreator, nil
}
