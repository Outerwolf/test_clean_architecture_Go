package application

import "go.uber.org/fx"

func ApplicationModule() fx.Option {
	return fx.Options(
		fx.Provide(),
	)
}
