package web_api_handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func ExampleHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// For any authenticated requests, we have to get our user first

		_, userId, _ := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)

		// if userId == 0{
		// 	http.Error(w, "There was an issue with your login session. Please logout and try logging in again.", 401)
		// 	return
		// }

		fmt.Println(userId)

		// define the parameters we want to accept at this API endpoint
		// THIS is where the parameters and their correct names and datatypes should be determined / assigned
		// NOT the frontend. The same is true for the response we send later.

		type RequestBody struct {
			PostInput string `json:"postInput"`
		}

		var requestBody RequestBody

		decodeErr := json.NewDecoder(r.Body).Decode(&requestBody)
		if decodeErr != nil {
			logger.Log(decodeErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		postInput := requestBody.PostInput

		createdExample, createdExampleErr := query_handlers.HandleCreateExampleQuery(postInput)
		if createdExampleErr != nil {
			http.Error(w, createdExampleErr.Message, createdExampleErr.Code)
			return
		}

		type Response struct {
			ExampleField string `json:"exampleField"`
		}

		response := Response{
			ExampleField: createdExample.ExampleField,
		}

		jsonRes, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			logger.Log(marshalErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(jsonRes)
	}
}
