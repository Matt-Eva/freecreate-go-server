package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GetMyWritingsParams struct {
	UserId int64
}

type CreatorWritingsGroup struct {
	CreatorName string
	CreatorUUID uuid.UUID
	Writing     []pg_core_types.MyWritingsStruct
}

func HandleGetMyWritings(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params GetMyWritingsParams) ([]CreatorWritingsGroup, *api_error.Error) {
	var creatorWritings []CreatorWritingsGroup

	getWritingParams := pg_core_types.GetMyWritingsParams{
		UserId: params.UserId,
	}

	myWritings, myWritingErr := pg_core_queries.GetMyWritings(ctx, pgCore, pgCoreQueries, getWritingParams)
	if myWritingErr != nil {
		return creatorWritings, myWritingErr
	}

	getCreatorsParams := pg_core_types.GetMyCreatorsParams{
		UserId: params.UserId,
	}

	myCreators, myCreatorsErr := pg_core_queries.GetMyCreators(ctx, pgCore, pgCoreQueries, getCreatorsParams)
	if myCreatorsErr != nil {
		return creatorWritings, myCreatorsErr
	}

	// creatorMap := make(map[int64]CreatorWritingsGroup)

	for i := 0; i < len(myCreators); i++{
		creatorGroup := CreatorWritingsGroup{
			CreatorName: myCreators[i].Name,
			CreatorUUID: myCreators[i].UUID,
		}

		// creatorMap[myCreators[i].ID] = creatorGroup

		for i := 0; i < len(myWritings); i++{
			writing := myWritings[i]
			creatorGroup.Writing = append(creatorGroup.Writing, writing)
		}

		creatorWritings = append(creatorWritings, creatorGroup)
	}

	// for i := 0; i < len(myWritings); i++{
	// 	writing := myWritings[i]

	// 	creator := creatorMap[writing.CreatorId]

	// 	creatorMap[writing.CreatorId].Writing = append(creator.Writing, writing)
	// }

	// for _, val := range creatorMap{
	// 	creatorWritings = append(creatorWritings, val)
	// }

	return creatorWritings, nil
}
