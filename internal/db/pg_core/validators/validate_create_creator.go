package pg_core_validators

import (
	"errors"
	"freecreate/internal/lib/api_error"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func ValidateCreator(args pgx.NamedArgs) *api_error.Error{
	if args["name"] == "" || args["name"] == nil {
		msg := "Name cannot be empty."
		err := errors.New(msg)
		apiErr := &api_error.Error{
			Code: http.StatusUnprocessableEntity,
			Message: msg,
			Error: err,
		}
		return apiErr
	}

	if args["user_id"] == 0 || args["user_id"] == nil {
		msg := "Invalid user ID."
		err := errors.New(msg)
		apiErr := &api_error.Error{
			Code: http.StatusUnprocessableEntity,
			Message: msg,
			Error: err,
		}
		return apiErr
	}

	return nil
}