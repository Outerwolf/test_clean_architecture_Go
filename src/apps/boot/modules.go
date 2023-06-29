package boot

import "go.uber.org/fx"

func CommonModules() fx.Option {
	return fx.Options(
		fx.Provide(NewRouter),
	)
}
