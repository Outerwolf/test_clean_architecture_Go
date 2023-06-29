package security

import (
	"github.com/auth/src/context/security/authentication/application"
	"github.com/auth/src/context/security/authentication/domain"
	"github.com/auth/src/context/security/authentication/infrastructure"
	"go.uber.org/fx"
)

func SecurityModule() fx.Option {
	return fx.Options(
		application.ApplicationModule(),
		domain.DomainModule(),
		infrastructure.InfrastructureModule(),
	)
}
