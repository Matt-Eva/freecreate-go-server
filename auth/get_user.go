package auth

import (
	"context"
	"errors"
	"fmt"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func GetUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) (session *sessions.Session, userId int, err *api_error.Error) {
	session, getSessionErr := sessionStore.Get(r, "user-session")
	if getSessionErr != nil {
		logger.Log(getSessionErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: getSessionErr,
		}

		return nil, 0, apiErr
	}

	uuidVal := session.Values["session_uuid"]
	if uuidVal == nil {
		return nil, 0, nil
	}

	sessionUuid, ok := uuidVal.(uuid.UUID)
	if !ok {
		err := errors.New("session uuid could not be converted to uuid")
		logger.Log(err)

		destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
		if destroySessionErr != nil {
			return nil, 0, destroySessionErr
		}

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: err,
		}

		return nil, 0, apiErr
	}

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)

	userIdString, getUserErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(authKey).Build()).ToString()
	if getUserErr != nil {
		logger.Log(getUserErr)
		
		destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
		if destroySessionErr != nil {
			return nil, 0, destroySessionErr
		}

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: getUserErr,
		}

		return nil, 0, apiErr
	}

	userId, parseIdErr := strconv.Atoi(userIdString)
	if parseIdErr != nil {
		logger.Log(parseIdErr)

		destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
		if destroySessionErr != nil {
			return nil, 0, destroySessionErr
		}

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: parseIdErr,
		}

		return nil, 0, apiErr
	}

	return session, userId, nil
}
