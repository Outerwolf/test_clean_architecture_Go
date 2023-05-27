package auth

import (
	"github.com/auth/src/auth/api"
	"github.com/auth/src/auth/application"
	"github.com/auth/src/auth/domain"
	"github.com/auth/src/auth/infrastructure"
	"go.uber.org/fx"
)

func AuthModule() fx.Option {
	return fx.Options(
		api.ApiModule(),
		application.ApplicationModule(),
		domain.DomainModule(),
		infrastructure.InfrastructureModule(),
	)
}
