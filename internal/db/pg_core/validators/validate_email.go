package pg_core_validators

import (
	"errors"
	"freecreate/internal/lib/api_error"
	"net/http"
)

func ValidateEmail(email string) *api_error.Error {
	if email == "" {
		msg := "Email cannot be empty."
		err := errors.New(msg)

		apiErr := &api_error.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: msg,
			Error:   err,
		}

		return apiErr
	}

	return nil
}
