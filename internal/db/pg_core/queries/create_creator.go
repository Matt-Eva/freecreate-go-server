package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateCreatorParams struct {
	UserId int64
	Name   string
	Handle string
}

type CreatedCreator struct {
	Name   string
	Handle string
	UUID   uuid.UUID
}

func CreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, createCreatorParams CreateCreatorParams) (CreatedCreator, *api_error.Error) {
	query := pgCoreQueries.CreateCreator()

	handle, handleErr := pg_core_validators.ValidateCreatorHandle(createCreatorParams.Handle)
	if handleErr != nil {
		return CreatedCreator{}, handleErr
	}

	namedArgs := pgx.NamedArgs{
		"name":           createCreatorParams.Name,
		"creator_handle": handle,
		"user_id":        createCreatorParams.UserId,
	}

	validateCreatorErr := pg_core_validators.ValidateCreator(namedArgs)
	if validateCreatorErr != nil {
		return CreatedCreator{}, validateCreatorErr
	}

	var name string
	var uuid uuid.UUID
	var creator_handle string

	rowResult := pgCore.QueryRow(ctx, query, namedArgs)

	createCreatorErr := rowResult.Scan(&name, &uuid, &creator_handle)

	var pgErr *pgconn.PgError
	if errors.As(createCreatorErr, &pgErr) && pgErr.Code == "23505" {
		apiErr := &api_error.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "That creator handle is already in use.",
			Error:   createCreatorErr,
		}

		return CreatedCreator{}, apiErr
	} else if createCreatorErr != nil {
		logger.Log(createCreatorErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   createCreatorErr,
		}

		return CreatedCreator{}, apiErr
	}

	createdCreator := CreatedCreator{
		Name:   name,
		Handle: creator_handle,
		UUID:   uuid,
	}

	return createdCreator, nil
}
