package auth

import (
	"context"
	"errors"
	"fmt"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/logger"
	"time"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func StoreOtp(ctx context.Context, valkeyClient valkey.Client, session *sessions.Session, email string, otp string) error {
	sessionUuid := session.Values["session_uuid"]
	if sessionUuid == nil {
		err := errors.New("Whoops, looks like your session id is invalid. Please try again.")
		logger.Log(err)
		return err
	}

	if len(otp) != 8 {
		err := errors.New("Sorry, that is not a valid One Time Password. One time password must be at least 8 characters.")
		logger.Log(err)
		return err
	}

	emailErr := pg_core_validators.ValidateEmail(email)
	if emailErr != nil {
		logger.Log(emailErr)
		return emailErr
	}

	otpKey := fmt.Sprintf("%s:%s", sessionUuid, email)

	// we don't have to check if the key value pair already exists, as set will overwrite it unless we call the NX function, which limits setting to only work if the record already does not exist.
	storeOtpErr := valkeyClient.Do(ctx, valkeyClient.B().Set().Key(otpKey).Value(otp).Ex(300 * time.Second).Build()).Error()
	if storeOtpErr != nil {
		logger.Log(storeOtpErr)
		return storeOtpErr
	}

	retrievedOtp, getOtpErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(otpKey).Build()).ToString()
	if getOtpErr != nil {
		logger.Log(getOtpErr)
		return getOtpErr
	}
	fmt.Println(retrievedOtp)

	return nil
}
