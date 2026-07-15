package auth

import (
	"errors"
	"freecreate/lib/logger"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func StoreOtp(valkeyClient valkey.Client,session *sessions.Session, email string, otp string)error{
	sessionUuid := session.Values["session_uuid"]
	if sessionUuid == nil {
		err := errors.New("Whoops, looks like your session id is invalid. Please try again.")
		logger.Log(err)
		return err
	}

	if len(otp) != 8{
		err := errors.New("Sorry, that is not a valid One Time Password. One time password must be at least 8 characters.")
		logger.Log(err)
		return err
	}	
	
	return nil
}