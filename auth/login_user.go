package auth

import (
	"context"
	"fmt"
	"freecreate/lib/logger"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LoginUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, r *http.Request, w http.ResponseWriter, userId int64) error {
	session, getSessionErr := sessionStore.Get(r, "user-session")
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		return getSessionErr
	}

	sessionUuid, newUuidErr := uuid.NewUUID()
	if newUuidErr != nil {
		logger.Log(newUuidErr)
		return newUuidErr
	}

	session.Values["session_uuid"] = sessionUuid

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)
	userIdString := fmt.Sprintf("%d", userId)

	storeUserErr := valkeyClient.Do(ctx, valkeyClient.B().Set().Key(authKey).Value(userIdString).Ex(12*time.Hour).Build()).Error()
	if storeUserErr != nil {
		logger.Log(storeUserErr)
		return storeUserErr
	}

	session.Save(r, w)

	return nil
}
