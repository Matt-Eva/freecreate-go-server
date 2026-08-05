package auth

import (
	"context"
	"fmt"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LoginUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, r *http.Request, w http.ResponseWriter, userId int64) *api_error.Error {
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

	sessionUuid, newUuidErr := uuid.NewUUID()
	if newUuidErr != nil {
		logger.Log(newUuidErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: newUuidErr,
		}

		return apiErr
	}

	session.Values["session_uuid"] = sessionUuid

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)
	userIdString := fmt.Sprintf("%d", userId)

	storeUserErr := valkeyClient.Do(ctx, valkeyClient.B().Set().Key(authKey).Value(userIdString).Ex(12*time.Hour).Build()).Error()
	if storeUserErr != nil {
		logger.Log(storeUserErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: storeUserErr,
		}

		return apiErr
	}

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
