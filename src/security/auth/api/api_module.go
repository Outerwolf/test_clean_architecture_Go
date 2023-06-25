package api

import (
	"github.com/auth/src/security/auth/api/auth_controller"
	"go.uber.org/fx"
)

func ApiModule() fx.Option {
	return fx.Options(
		fx.Provide(
			auth_controller.NewAuthController,
			auth_controller.NewAuthRoutes,
		),
	)
}
