package api

import (
	http1_auth "github.com/auth/src/auth/api/auth_http1"
	"go.uber.org/fx"
)

func ApiModule() fx.Option {
	return fx.Options(
		fx.Invoke(http1_auth.RegisterHttpEndpoints),
	)
}
