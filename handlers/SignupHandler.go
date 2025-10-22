package handlers

import (
	"encoding/json"
	"freecreate/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)


func SignupHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		type Body struct {
			Email string `json:"email"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
		}

		// email := body.Email


		// need to check for valid email address before creating user

		// var currentUser config.User;
		// var newUser config.User;

		// result := gormPGClient.Where("email = ?", email).First(&currentUser)
		// if result.Error != nil && result.Error == gorm.ErrRecordNotFound{
		// 	newUser.Email = email;
		// 	result := gormPGClient.Create(&newUser);
		// 	if result.Error != nil {
		// 		logger.Log(result.Error)
		// 		http.Error(w, result.Error.Error(), http.StatusUnprocessableEntity)
		// 		return
		// 	}
		// } else {
		// 	err := errors.New("email address already in use")
		// 	logger.Log(err)
		// 	http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		// 	return 
		// }

		// session, _ := sessionStore.Get(r, "user-session")
		// session.Values["userId"] = newUser.ID
		// err := session.Save(r, w)
		// if err != nil {
		// 	logger.Log(err)
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	}
}