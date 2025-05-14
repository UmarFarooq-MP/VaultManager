package main

import (
	"fmt"
	"vault/internal/provider"
	"vault/internal/secrets"
	"vault/internal/shared"
)

func main() {
	var secretProvider secrets.SecretsProvider
	appConfig := shared.LoadAppConfig()
	secretProvider, err := provider.NewHashicrop(appConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(secretProvider.GetAllSecrets())
}
