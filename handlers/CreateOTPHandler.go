package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"freecreate/logger"
	"math/big"
	"net/http"
	"strconv"

	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func CreateOTPHandler(resendClient *resend.Client, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Body struct {
			Email string `json:"email"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)

			errorResponse := ErrorResponse{
				Message: jErr.Error(),
			}

			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		var otp string
		for i := 0; i < 8; i++ {
			a, _ := rand.Int(rand.Reader, big.NewInt(10))
			str := strconv.FormatInt(a.Int64(), 10)
			otp += str
		}

		email := body.Email

		html := fmt.Sprintf("<p>Here is your FreeCreate One Time Password: %s</p><p>This password will expire after 5 minutes</p><p>DO NOT share this with anyone. WE WILL NEVER ASK YOU FOR YOUR ONE TIME PASSWORD </p>", otp)

		fmt.Println(html)
 
		params := &resend.SendEmailRequest{
			From:    "test@email.freecreate.net",
			To:      []string{email},
			Html:    html,
			Subject: "Hello from FreeCreate!",
		}

		_, rErr := resendClient.Emails.Send(params)
		if rErr != nil {
			logger.Log(rErr)

			errorResponse := ErrorResponse {
				Message : rErr.Error(),
			}	

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}


		type Response struct {
			Data string `json:"data"`
		}

		response := Response{
			Data: "success",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
