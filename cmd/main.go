package main

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"log"
)

func main() {

	config := vault.DefaultConfig()
	config.Address = "http://127.0.0.1:8800"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("failed to create Vault client: %v", err)
	}

	client.SetToken("set_the_token")

	ctx := context.Background()
	secret, err := client.KVv2("secret").Get(ctx, "staging")
	if err != nil {
		log.Fatalf("failed to read secret: %v", err)
	}

	for key, val := range secret.Data {
		strVal := fmt.Sprintf("%v", val)
		fmt.Printf("%s=%s\n", key, strVal)
	}
}
