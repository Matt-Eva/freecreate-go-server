package pg_core_validators

import (
	"errors"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func ValidateUserInfo(args pgx.NamedArgs) *api_error.Error {
	apiErr := &api_error.Error{
		Code: http.StatusUnprocessableEntity,
	}

	if args["user_id"] == 0 {
		err := errors.New("user_id is empty")
		logger.Log(err)
		apiErr.Code = http.StatusInternalServerError
		apiErr.Message = api_error.InteralServerErrorMessage
		apiErr.Error = err
		return apiErr
	}

	return nil
}
