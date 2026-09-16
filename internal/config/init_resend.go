package config

import (
	"fmt"
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

func InitResend() *resend.Client {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		log.Fatal("resend api Key cannot be empty")
	}
	client := resend.NewClient(apiKey)

	fmt.Println("resend connection successful!")

	return client
}
