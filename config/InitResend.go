package config

import (
	"os"

	"github.com/resend/resend-go/v2"
)

func InitResend() *resend.Client {
	apiKey := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(apiKey)
	return client
}
