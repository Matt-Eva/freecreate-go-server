package handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/logger"
	"net/http"

	"github.com/resend/resend-go/v2"
)

func EmailHandler(resendClient *resend.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Body struct {
			Email string
		}

		var body Body

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := body.Email

		params := &resend.SendEmailRequest{
			From: "test@email.freecreate.net",
			To: []string{email},
			Html: "<p>Hello from FreeCreate!</p>",
			Subject: "Hello from FreeCreate!",
		}

		sent, err := resendClient.Emails.Send(params)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(sent.Id)
	}
}
