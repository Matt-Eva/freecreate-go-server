package email_handler

import (
	"fmt"
	"freecreate/lib/logger"

	"github.com/resend/resend-go/v2"
)

func SendOtp(resendClient *resend.Client, email string, otp string) error {
	html := fmt.Sprintf("<div><p>Your One Time Password is %s</p> <p>This password will expire in 5 minutes.</p> <p>DO NOT SHARE your password with ANYONE. We will NEVER ask you to share your one time password with us.</p> </div>", otp)

	params := &resend.SendEmailRequest{
		From:    "test@email.freecreate.net",
		To:      []string{email},
		Html:    html,
		Subject: "FreeCreate One Time Password",
	}

	_, sendEmailErr := resendClient.Emails.Send(params)
	if sendEmailErr != nil {
		fmt.Println(sendEmailErr.Error())
		logger.Log(sendEmailErr)
		return sendEmailErr
	}

	return nil
}
