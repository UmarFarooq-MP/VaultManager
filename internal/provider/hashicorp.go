package provider

import (
	"context"
	"fmt"

	vault "github.com/hashicorp/vault/api"
	"vault/internal/secrets"
	"vault/internal/shared"
)

type HashiCorpConfig struct {
	SecretKey    string `json:"secret_key"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type Hashicrop struct {
	kv     *vault.KVv2
	config *shared.AppConfig
}

func NewHashicrop(config *shared.AppConfig) (*Hashicrop, error) {
	cfg := vault.DefaultConfig()
	cfg.Address = config.ProviderConfig.Addr

	// sincee it's a one time operation so we can create a client here
	client, err := vault.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	client.SetToken(config.ProviderConfig.Token)

	return &Hashicrop{
		kv:     client.KVv2(config.ProviderConfig.MountPath),
		config: config,
	}, nil
}

func (v *Hashicrop) GetAllSecrets() (map[string]string, error) {
	path := v.config.ProviderConfig.SecretPath
	secret, err := v.kv.Get(context.Background(), path)
	if err != nil {
		return nil, err
	}
	out := make(map[string]string)
	for k, v := range secret.Data {
		out[k] = fmt.Sprintf("%v", v)
	}
	return out, nil
}

// Assert interface
var _ secrets.SecretsProvider = (*Hashicrop)(nil)
