package secrets

type SecretsProvider interface {
	GetAllSecrets() (map[string]string, error)
}
