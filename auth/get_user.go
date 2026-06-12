package auth

import (
	"freecreate/pgModels"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func GetUser(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB, w http.ResponseWriter, r *http.Request) (uint, error) {
	sessionUUID, aErr := CheckSession(sessionStore, w, r)
	if aErr != nil {
		return 0, aErr
	}

	var userId uint

	uErr := gormPGClient.Model(pgModels.User{}).Select("id").Where("session_uuid = ?", sessionUUID).First(&userId).Error
	if uErr != nil {
		return 0, uErr
	}

	return userId, nil
}
