package config

import (
	"go.uber.org/fx"
	"vault/internal/shared/config/environment"
)

var Module = fx.Module(
	"configfx",
	fx.Provide(func() environment.Environment {
		return environment.ConfigAppEnv()
	},
	))

var ModuleFunc = func(e environment.Environment) fx.Option {
	return fx.Module(
		"configfx",
		fx.Provide(func() environment.Environment {
			return environment.ConfigAppEnv(e)
		}),
	)
}
