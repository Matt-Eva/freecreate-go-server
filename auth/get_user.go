package auth

import (
	"context"
	"errors"
	"fmt"
	"freecreate/lib/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func GetUser(ctx context.Context, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, w http.ResponseWriter, r *http.Request) (session *sessions.Session, userId int, error error) {
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

	authKey := fmt.Sprintf("auth_key:%s", sessionUuid)

	userIdString, getUserErr := valkeyClient.Do(ctx, valkeyClient.B().Get().Key(authKey).Build()).ToString()
	if getUserErr != nil {
		logger.Log(getUserErr)
		err := errors.New("could not retrive user id")
		DestroySession(sessionStore, w, r)
		return nil, 0, err
	}

	userId, parseIdErr := strconv.Atoi(userIdString)
	if parseIdErr != nil {
		err := errors.New("could not convert userid string to int64")
		logger.Log(err)
		
		return nil, 0, err
	}

	return session, userId, nil
}
