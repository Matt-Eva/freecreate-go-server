package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreatedCreator struct{
	Name string
	UUID uuid.UUID
}

func CreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, creatorName string, userId int64)(CreatedCreator, *api_error.Error){
	query := pgCoreQueries.CreateCreator()
	namedArgs := pgx.NamedArgs{
		"name": creatorName,
		"user_id": userId,
	}

	validateCreatorErr := pg_core_validators.ValidateCreator(namedArgs)
	if validateCreatorErr != nil {
		return newCreator{}, validateCreatorErr
	}

	var name string
	var uuid uuid.UUID

	createCreatorErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&name, &uuid)
	if createCreatorErr != nil {
		logger.Log(createCreatorErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: createCreatorErr,
		}

		return newCreator{}, apiErr
	}

	newCreator := newCreator{
		Name: name,
		UUID: uuid,
	}

	return newCreator, nil
}