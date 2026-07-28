package auth

import (
	"context"
	"errors"
	"fmt"
	"freecreate/lib/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func GetUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) (session *sessions.Session, userId int64, error error) {
	isLoggedIn, loggedInErr := CheckLogin(sessionStore, w, r)
	if loggedInErr != nil {
		logger.Log(loggedInErr)
		return nil, 0, loggedInErr
	}
	
	if !isLoggedIn {
		err := errors.New("session cookie does not have logged in attribute set to true")
		logger.Log(err)
		return nil, 0, err
	}
	
	session, sessionUuid, getSessionErr := GetSession(sessionStore, w, r)
	if getSessionErr != nil {
		logger.Log(getSessionErr)
		return nil, 0, getSessionErr
	}

	authKey := fmt.Sprintf("auth_key::%s", sessionUuid)

	userId, getUserErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(authKey).Build()).ToInt64()
	if getUserErr != nil {
		logger.Log(getUserErr)
		err := errors.New("could not retrive user id")
		return nil, 0, err
	}

	return session, userId, nil
}
