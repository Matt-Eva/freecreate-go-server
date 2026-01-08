package handlers

import (
	"fmt"
	"freecreate/auth"
	"freecreate/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func UpdateWritingHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		userId, uErr := auth.GetUser(sessionStore, gormPGClient, w, r)
		if uErr != nil {
			logger.Log(uErr)
			http.Error(w, uErr.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(userId)

		type Body struct {
			
		}
	}
}