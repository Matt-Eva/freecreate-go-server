package web_auth

import (
	"context"
	"errors"
	"fmt"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/valkey-io/valkey-go"
)

func ValidateOtp(ctx context.Context, sessionUuid uuid.UUID, valkeyClient valkey.Client, email string, otp string) *api_error.Error {
	emailValidationErr := pg_core_validators.ValidateEmail(email)
	if emailValidationErr != nil {
		return emailValidationErr
	}

	if otp == "" {
		msg := "otp cannot be empty"
		err := errors.New(msg)

		apiErr := &api_error.Error{
			Code: http.StatusUnprocessableEntity,
			Message: msg,
			Error: err,
		}

		return apiErr
	} else if len(otp) != 8 {
		msg := "otp is invalid length"
		err := errors.New(msg)

		apiErr := &api_error.Error{
			Code: http.StatusUnprocessableEntity,
			Message: msg,
			Error: err,
		}

		return apiErr
	}

	otpKey := fmt.Sprintf("%s:%s", sessionUuid, email)

	retrievedOtp, getOtpErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(otpKey).Build()).ToString()
	if getOtpErr != nil {
		logger.Log(getOtpErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: getOtpErr,
		}

		return apiErr
	}

	if retrievedOtp != otp {
		msg := "Either your email or one time password are invalid."
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
