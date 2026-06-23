package config

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

func InitResend() *resend.Client {
	apiKey := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(apiKey)

	fmt.Println("resend connection successful!")

	return client
}
