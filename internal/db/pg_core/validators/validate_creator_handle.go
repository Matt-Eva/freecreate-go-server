package pg_core_validators

import (
	"errors"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"
	"strings"
)

func ValidateCreatorHandle(handle string)(string, *api_error.Error){
	 apiErr := &api_error.Error{
		Code: http.StatusInternalServerError,
		Message: api_error.InteralServerErrorMessage,
	}

	if handle == ""{
		err := errors.New("handle cannot be empty")
		logger.Log(err)
		apiErr.Error = err
		return "", apiErr
	}

	validatedHandle := "@" + strings.ReplaceAll(handle, " ", "-")
	
	if validatedHandle == "@" || validatedHandle == ""{
		err := errors.New("validated handle cannot be empty")
		logger.Log(err)
		apiErr.Error = err
		return "", apiErr
	}

	return validatedHandle, nil
}