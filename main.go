package main

import (
	"github.com/auth/src/boot"
	"github.com/auth/src/boot/routes"
	"github.com/auth/src/config"
	"github.com/auth/src/security"
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
		),
		boot.CommonModules(),
		security.SecurityModule(),

		fx.Invoke(
			routes.SetRoutes,
			boot.StartHttpServer,
		),
	)
}
