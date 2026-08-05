package auth

import (
	"context"
	"errors"
	"fmt"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func DestroyUserSession(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) *api_error.Error {
	session, getSessionErr := sessionStore.Get(r, "user-session")
	if getSessionErr != nil {
		logger.Log(getSessionErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: getSessionErr,
		}

		return apiErr
	}

	var errs []error

	sessionUuid := session.Values["session_uuid"]

	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1

	saveSessionErr := session.Save(r, w)
	if saveSessionErr != nil {
		logger.Log(saveSessionErr)

		errs = append(errs, saveSessionErr)
	}

	if sessionUuid != nil {
		authKey := fmt.Sprintf("auth_key:%s", sessionUuid)

		deleteRecordErr := valkeyClient.Do(ctx, valkeyClient.B().Del().Key(authKey).Build()).Error()
		if deleteRecordErr != nil {
			logger.Log(deleteRecordErr)

			errs = append(errs, deleteRecordErr)
		}
	}

	finalErrs := errors.Join(errs...)
	if finalErrs != nil {
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: finalErrs,
		}
		
		return apiErr
	}
	
	return nil
}

func DestroyGuestSession(sessionStore *sessions.CookieStore, w http.ResponseWriter, r *http.Request) *api_error.Error {
	session, getSessionErr := sessionStore.Get(r, "user-session")
	if getSessionErr != nil {
		logger.Log(getSessionErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: getSessionErr,
		}

		return apiErr
	}

	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1

	saveSessionErr := session.Save(r, w)
	if saveSessionErr != nil {
		logger.Log(saveSessionErr)

		apiErr := &api_error.Error{
				Code: http.StatusInternalServerError,
				Message: api_error.InteralServerErrorMessage,
				Error: saveSessionErr,
			}

		return apiErr
	}

	return nil
}
