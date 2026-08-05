package web_auth

import (
	"context"
	"errors"
	"fmt"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/valkey-io/valkey-go"
)

func StoreOtp(ctx context.Context, valkeyClient valkey.Client, sessionUuid uuid.UUID, email string, otp string) *api_error.Error {

	if len(otp) != 8 {
		msg := "Sorry, that is not a valid One Time Password. One time password must be at least 8 characters."
		err := errors.New(msg)

		apiErr := &api_error.Error{
			Code: http.StatusUnprocessableEntity,
			Message: msg,
			Error: err,
		}

		return apiErr
	}

	emailErr := pg_core_validators.ValidateEmail(email)
	if emailErr != nil {
		return emailErr
	}

	otpKey := fmt.Sprintf("%s:%s", sessionUuid, email)

	// we don't have to check if the key value pair already exists, as set will overwrite it unless we call the NX function, which limits setting to only work if the record already does not exist.
	storeOtpErr := valkeyClient.Do(ctx, valkeyClient.B().Set().Key(otpKey).Value(otp).Ex(12*time.Hour).Build()).Error()
	if storeOtpErr != nil {
		logger.Log(storeOtpErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: storeOtpErr,
		}

		return apiErr
	}

	return nil
}
