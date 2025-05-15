package config

import (
	"github.com/sarulabs/di"

	"vault/internal/shared/config/environment"
)

func AddEnv(container *di.Builder) error {
	envDep := di.Def{
		Name:  "env",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return environment.ConfigAppEnv(), nil
		},
	}
	return container.Add(envDep)
}
