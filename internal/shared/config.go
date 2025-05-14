package shared

import (
	"log"
	"os"
)

type AppConfig struct {
	Provider       string
	ProviderConfig *hcConfig
}

type hcConfig struct {
	Addr       string
	Token      string
	MountPath  string
	SecretPath string
}

func LoadAppConfig() *AppConfig {
	provider := getEnv("SECRETS_PROVIDER", "hashicorp")
	cfg := &AppConfig{
		Provider: provider,
	}

	if provider == "hashicorp" {
		cfg.ProviderConfig = &hcConfig{
			Addr:       getEnv("VAULT_ADDR", "http://127.0.0.1:8800"),
			Token:      getEnv("VAULT_TOKEN", ""),
			MountPath:  getEnv("VAULT_MOUNT", "secret"),
			SecretPath: getEnv("VAULT_SECRET_PATH", "staging"),
		}
		if cfg.ProviderConfig.Token == "" {
			log.Fatal("VAULT_TOKEN is required")
		}
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
