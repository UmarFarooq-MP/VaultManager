package config

import (
	"encoding/json"
	"fmt"

	"github.com/caarlos0/env/v8"
	"github.com/mcuadros/go-defaults"
	"vault/internal/secrets"
	"vault/internal/shared/config/environment"

	typeMapper "vault/internal/reflection/type_mapper"
)

func BindConfig[T any](provider *secrets.SecretsProvider, environments ...environment.Environment) (T, error) {
	return BindConfigKey[T]("", provider, environments...)
}

func BindConfigKey[T any](configKey string, provider *secrets.SecretsProvider, environments ...environment.Environment) (T, error) {

	var zero T

	if provider == nil || *provider == nil {
		return zero, fmt.Errorf("secrets provider is nil")
	}

	cfg := typeMapper.GenericInstanceByT[T]()

	secretsMap, err := (*provider).GetAllSecrets()
	if err != nil {
		return zero, fmt.Errorf("failed to get secrets: %w", err)
	}

	if configKey == "" {
		return zero, fmt.Errorf("config key must be provided")
	}

	rawJSON, ok := secretsMap[configKey]
	if !ok {
		return zero, fmt.Errorf("config key %q not found in Vault secret", configKey)
	}

	// Parse stringified JSON value
	var rawData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJSON), &rawData); err != nil {
		return zero, fmt.Errorf("failed to unmarshal raw secret: %w", err)
	}

	// Re-marshal for decoding into struct
	jsonBytes, err := json.Marshal(rawData)
	if err != nil {
		return zero, fmt.Errorf("failed to re-marshal for decoding: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, cfg); err != nil {
		return zero, fmt.Errorf("failed to unmarshal into struct: %w", err)
	}

	_ = env.Parse(cfg)
	defaults.SetDefaults(cfg)

	return cfg, nil
}
