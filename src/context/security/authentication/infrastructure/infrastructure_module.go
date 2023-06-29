package infrastructure

import (
	"github.com/auth/src/context/security/authentication/infrastructure/repository"
	"go.uber.org/fx"
)

func InfrastructureModule() fx.Option {
	return fx.Options(
		fx.Provide(
			repository.NewMongoAuthRepository,
		))
}
