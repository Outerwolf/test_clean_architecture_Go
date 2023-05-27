package main

import (
	"github.com/auth/src/auth"
	"github.com/auth/src/boot"
	"github.com/auth/src/config"
	"go.uber.org/fx"
)

func main() {
	println("Hello, World!")
	newFxApp().Run()
}

func newFxApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfigBuilder,
			config.NewConfiguration,
			boot.CreateHTTPGinServer,
		),
		auth.AuthModule(),

		//
		fx.Invoke(
			boot.RegisterHealthCheck,
			boot.StartHttpGinServer,
		),
	)
}
