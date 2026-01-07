package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func UpdateWritingHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){}
}