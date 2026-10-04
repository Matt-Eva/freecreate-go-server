package pg_core_validators

import (
	"errors"
	"freecreate/internal/lib/api_error"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func ValidateNewWriting(args pgx.NamedArgs) *api_error.Error {
	apiErr := &api_error.Error{
		Code: http.StatusUnprocessableEntity,
	}

	if args["creator_id"] == 0 {
		msg := "creator id cannot be empty"
		err := errors.New(msg)
		apiErr.Message = msg
		apiErr.Error = err
		return apiErr
	}

	if args["user_id"] == 0 {
		msg := "user id cannot be empty"
		err := errors.New(msg)
		apiErr.Message = msg
		apiErr.Error = err
		return apiErr
	}

	if args["writing_type"] == "" || args["writing_type"] == nil {
		msg := "writing type cannot be empty"
		err := errors.New(msg)
		apiErr.Message = msg
		apiErr.Error = err
		return apiErr
	}

	return nil
}
