package web_auth

import (
	"context"
	"errors"
	"fmt"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func GetUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) (sessionUuid uuid.UUID, userId int64, err *api_error.Error) {
	session, getSessionErr := sessionStore.Get(r, "user-session")
	if getSessionErr != nil {
		logger.Log(getSessionErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   getSessionErr,
		}

		return uuid.Nil, 0, apiErr
	}

	uuidVal := session.Values["session_uuid"]
	if uuidVal == nil {
		return uuid.Nil, 0, nil
	}

	sessionUuid, ok := uuidVal.(uuid.UUID)
	if !ok {
		err := errors.New("session uuid could not be converted to uuid")
		logger.Log(err)

		destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
		if destroySessionErr != nil {
			return uuid.Nil, 0, destroySessionErr
		}

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}

		return uuid.Nil, 0, apiErr
	}

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)

	userIdString, getUserErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(authKey).Build()).ToString()
	if getUserErr != nil {
		logger.Log(getUserErr)

		destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
		if destroySessionErr != nil {
			return uuid.Nil, 0, destroySessionErr
		}

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   getUserErr,
		}

		return uuid.Nil, 0, apiErr
	}

	userId, parseIdErr := strconv.ParseInt(userIdString, 10, 64)
	if parseIdErr != nil {
		logger.Log(parseIdErr)

		destroySessionErr := DestroyUserSession(ctx, sessionStore, valkeyClient, w, r)
		if destroySessionErr != nil {
			return uuid.Nil, 0, destroySessionErr
		}

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   parseIdErr,
		}

		return uuid.Nil, 0, apiErr
	}

	return sessionUuid, userId, nil
}
