package web_page_utils

import (
	"errors"
	"freecreate/logger"
	"net/http"
)

func GetFormAction(r *http.Request)(string, error){
	parseFormErr := r.ParseForm()
	if parseFormErr != nil {
		logger.Log(parseFormErr)
		return "", parseFormErr
	}

	formAction := r.FormValue("form_action")
	if formAction == ""{
		err := errors.New("form action is empty string")
		logger.Log(err)
		return "", err
	}

	return formAction, nil
}