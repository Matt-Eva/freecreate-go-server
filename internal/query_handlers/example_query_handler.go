package query_handlers

import (
	"context"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

// ========= Query Handler Return Values =========

// ========== Final Return Value
type ExampleQueryReturnValues struct {
	ExampleParams []ExampleQueryReturnValue
}

// ========== Child Return Values

type ExampleQueryReturnValue struct {
	ExampleName string
	ExampleUUID uuid.UUID
}

// ========= Query Handler Logic ===================

func HandleExampleQuery(ctx context.Context, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) (ExampleQueryReturnValues, *api_error.Error) {
	var returnValues ExampleQueryReturnValues

	for i := 0; i < 10; i++ {
		returnValue := ExampleQueryReturnValue{
			ExampleName: "Example",
			ExampleUUID: uuid.New(),
		}

		returnValues.ExampleParams = append(returnValues.ExampleParams, returnValue)
	}

	return returnValues, nil
}
