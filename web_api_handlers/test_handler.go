package web_api_handlers

import (
	"errors"
	"freecreate/logger"
	"net/http"
)

func TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errMsg := "I want to generate this error message by default"
		err := errors.New(errMsg)
		logger.Log(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
