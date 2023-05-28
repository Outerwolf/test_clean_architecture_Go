package application

import (
	sign_up "github.com/auth/src/security/auth/application/usecase/sign-up"
	"go.uber.org/fx"
)

func ApplicationModule() fx.Option {
	return fx.Options(
		fx.Provide(
			sign_up.NewSignUpUseCase,
		),
	)
}
