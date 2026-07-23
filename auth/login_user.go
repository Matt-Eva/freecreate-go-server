package auth

import (
	"context"
	"fmt"
	"freecreate/lib/logger"
	"time"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LoginUser(ctx context.Context, session *sessions.Session, valkeyClient valkey.Client, userId int) error {
	session.Values["logged_in"] = true
	sessionUuid := session.Values["session_uuid"]

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)
	userIdString := fmt.Sprintf("%d", userId)

	storeOtpErr := valkeyClient.Do(ctx, valkeyClient.B().Set().Key(authKey).Value(userIdString).Ex(300*time.Second).Build()).Error()
	if storeOtpErr != nil {
		logger.Log(storeOtpErr)
		return storeOtpErr
	}

	return nil
}
