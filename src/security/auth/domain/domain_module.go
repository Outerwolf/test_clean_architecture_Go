package domain

import "go.uber.org/fx"

func DomainModule() fx.Option {
	return fx.Options(
		fx.Provide(),
	)
}
