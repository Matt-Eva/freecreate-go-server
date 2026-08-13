package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MyCreatorsStruct struct {
	Name string
	UUID uuid.UUID
}

func GetMyCreators(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, userId int64) ([]MyCreatorsStruct, *api_error.Error) {
	query := pgCoreQueries.GetMyCreators()
	namedArgs := pgx.NamedArgs{
		"user_id": userId,
	}

	queryResult, queryErr := pgCore.Query(ctx, query, namedArgs)

	if errors.Is(queryErr, pgx.ErrNoRows){
		return []MyCreatorsStruct{}, nil
	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}
		return []MyCreatorsStruct{}, apiErr
	}

	var myCreators []MyCreatorsStruct

	for queryResult.Next(){		
		var myCreator MyCreatorsStruct
		scanErr := queryResult.Scan(&myCreator)
		if scanErr != nil {
			logger.Log(scanErr)
		}
		myCreators = append(myCreators, myCreator)
	}
	
	return myCreators, nil
}
