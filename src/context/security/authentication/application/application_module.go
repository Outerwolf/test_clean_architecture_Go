package application

import (
	"github.com/auth/src/context/security/authentication/application/usecase/sign-up"
	"go.uber.org/fx"
)

func ApplicationModule() fx.Option {
	return fx.Options(
		fx.Provide(
			sign_up.NewSignUpUseCase,
		),
	)
}
