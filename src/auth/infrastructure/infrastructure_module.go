package infrastructure

import "go.uber.org/fx"

func InfrastructureModule() fx.Option {
	return fx.Options(
		fx.Provide(),
	)
}
