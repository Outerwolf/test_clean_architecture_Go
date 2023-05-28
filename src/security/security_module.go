package security

import (
	"github.com/auth/src/security/auth/api"
	"github.com/auth/src/security/auth/application"
	"github.com/auth/src/security/auth/domain"
	"github.com/auth/src/security/auth/infrastructure"
	"go.uber.org/fx"
)

func SecurityModule() fx.Option {
	return fx.Options(
		api.ApiModule(),
		application.ApplicationModule(),
		domain.DomainModule(),
		infrastructure.InfrastructureModule(),
	)
}
