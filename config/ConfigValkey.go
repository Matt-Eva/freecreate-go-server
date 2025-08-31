package config

import (
	"fmt"
	"log"

	"github.com/valkey-io/valkey-go"
)

func ConfigValkey() valkey.Client {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		log.Fatalf("could not connect to valkey: %W", err)
	}
	fmt.Println("connection to valkey successful!")

	return client
}
