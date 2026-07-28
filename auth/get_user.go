package auth

import (
	"context"
	"errors"
	"fmt"
	"freecreate/lib/logger"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func GetUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) (session *sessions.Session, userId int, error error) {
	session, getSessionErr := sessionStore.Get(r, "user-session")
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		return nil, 0, getSessionErr
	}

	uuidVal := session.Values["session_uuid"]
	if uuidVal == nil {
		return nil, 0, nil
	}

	sessionUuid, ok := uuidVal.(uuid.UUID)
	if !ok {
		err := errors.New("session uuid could not be converted to uuid")
		logger.Log(err)
		DestroyUserSession(session, w, r)
		return nil, 0, err
	}

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)

	userIdString, getUserErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(authKey).Build()).ToString()
	if getUserErr != nil {
		logger.Log(getUserErr)
		err := errors.New("could not retrive user id")
		DestroyUserSession(session, w, r)
		return nil, 0, err
	}

	userId, parseIdErr := strconv.Atoi(userIdString)
	if parseIdErr != nil {
		err := errors.New("could not convert userid string to int64")
		logger.Log(err)
		DestroyUserSession(session, w, r)
		return nil, 0, err
	}

	return session, userId, nil
}
