package controller

import (
	"github.com/auth/src/apps/controller/authentication"
	"go.uber.org/fx"
)

func ApiModule() fx.Option {
	return fx.Options(
		fx.Provide(
			authentication.NewAuthController,
			authentication.NewAuthRoutes,
		),
	)
}
