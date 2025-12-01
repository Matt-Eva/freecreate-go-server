package handlers

import (
	"fmt"
	"freecreate/auth"
	"freecreate/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func DeleteCreatorHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc{
return func (w http.ResponseWriter, r *http.Request){
	userId, aErr := auth.CheckSession(sessionStore, w, r)
	if aErr != nil {
		logger.Log(aErr)
		http.Error(w, aErr.Error(), http.StatusUnauthorized)
		return
	}

	creatorId := chi.URLParam(r, "creatorId")
	fmt.Println("userId", userId)
	fmt.Println("creatorId", creatorId)
}
}