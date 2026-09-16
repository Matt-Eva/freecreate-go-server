package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/lib/api_error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GetMyCreatorsParams struct {
	UserId int64
}

type MyCreatorsStruct struct {
	Name          string
	CreatorHandle string
	UUID          uuid.UUID
}

func HandleGetMyCreators(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, getMyCreatorsParams GetMyCreatorsParams) ([]MyCreatorsStruct, *api_error.Error) {

	var myCreators []MyCreatorsStruct

	creators, getMyCreatorsErr := pg_core_queries.GetMyCreators(ctx, pgCore, pgCoreQueries, getMyCreatorsParams.UserId)
	if getMyCreatorsErr != nil {
		return myCreators, getMyCreatorsErr
	}

	for i := 0; i < len(creators); i++ {
		creator := creators[i]
		myCreator := MyCreatorsStruct{
			Name:          creator.Name,
			CreatorHandle: creator.CreatorHandle,
			UUID:          creator.UUID,
		}

		myCreators = append(myCreators, myCreator)
	}

	return myCreators, nil
}
