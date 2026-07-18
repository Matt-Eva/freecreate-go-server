package web_api_handlers

import (
	"net/http"
)

func SignupRequestOtp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "there was an error requesting your one time password", 422)

		// type ResponsePayload struct {
		// 	Message string `json:"message"`
		// }

		// responsePayload := ResponsePayload{
		// 	Message: "success!",
		// }

		// response, marshalError := json.Marshal(responsePayload)
		// if marshalError != nil {
		// 	logger.Log(marshalError)
		// 	http.Error(w, "error marshalling json response", 500)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusAccepted)
		// w.Write(response)
	}
}
