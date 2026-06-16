package lib

import (
	"errors"
	"fmt"
	"net/http"
)

func HandleInvalidWebpageRequestMethod(w http.ResponseWriter, requestMethod string){
	errMsg := fmt.Sprintf("not a valid request method: method %s is not GET or POST", requestMethod)
	err := errors.New(errMsg)
	http.Error(w, err.Error(), http.StatusBadRequest)
}