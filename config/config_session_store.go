package config

import (
	"encoding/base64"
	"freecreate/lib/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)


func ConfigSessionStore(environment string)(*sessions.CookieStore, error){

	sessionAuthKey, err := base64.StdEncoding.DecodeString(os.Getenv("SESSION_AUTH_KEY"))
		if err != nil {
			logger.Log(err)
			log.Fatal(err.Error())
			return nil, err
		}
	
		sessionEncryptionKey, err := base64.StdEncoding.DecodeString(os.Getenv("SESSION_ENCRYPTION_KEY"))
		if err != nil {
			logger.Log(err)
			log.Fatal(err.Error())
			return nil, err
		}
	
		sessionStore := sessions.NewCookieStore(sessionAuthKey, sessionEncryptionKey)
	
		if environment == "PRODUCTION" {
			sessionStore.Options = &sessions.Options{
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
				HttpOnly: true,
			}
		}
		return sessionStore, nil
}