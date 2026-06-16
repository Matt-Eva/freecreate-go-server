package lib

import (
	"errors"
	"fmt"
	"freecreate/logger"
	"net/http"
)

func HandleInvalidWebpageRequestMethod(w http.ResponseWriter, requestMethod string) {
	errMsg := fmt.Sprintf("not a valid request method: method %s is not GET or POST", requestMethod)
	err := errors.New(errMsg)
	logger.Log(err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}
