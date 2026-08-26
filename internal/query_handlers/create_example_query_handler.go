package query_handlers

import (
	"errors"
	"freecreate/internal/lib/api_error"
	"net/http"
)

type CreatedExample struct {
	ExampleField string
}

func HandleCreateExampleQuery(postInput string) (CreatedExample, *api_error.Error) {

	var createdExample CreatedExample

	if postInput == "" {
		msg := "Post input cannot be empty."
		err := errors.New(msg)
		apiErr := &api_error.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: msg,
			Error:   err,
		}

		return createdExample, apiErr
	}

	createdExample.ExampleField = postInput

	return createdExample, nil

}
