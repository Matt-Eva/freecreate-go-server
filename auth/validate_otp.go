package auth

import (
	"context"
	"errors"
	"fmt"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/logger"

	"github.com/google/uuid"
	"github.com/valkey-io/valkey-go"
)

func ValidateOtp(ctx context.Context, sessionUuid uuid.UUID, valkeyClient valkey.Client, email string, otp string) error {
	emailValidationErr := pg_core_validators.ValidateEmail(email)
	if emailValidationErr != nil {
		logger.Log(emailValidationErr)
		return emailValidationErr
	}

	if otp == "" {
		err := errors.New("otp cannot be empty")
		logger.Log(err)
		return err
	} else if len(otp) != 8 {
		err := errors.New("otp is invalid length")
		return err
	}

	otpKey := fmt.Sprintf("%s:%s", sessionUuid, email)
	retrievedOtp, getOtpErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(otpKey).Build()).ToString()
	if getOtpErr != nil {
		logger.Log(getOtpErr)
		return getOtpErr
	}

	if retrievedOtp != otp {
		err := errors.New("Either your email or one time password are invalid.")
		logger.Log(err)
		return err
	}

	return nil
}
