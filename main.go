package main

import (
	"github.com/auth/src/apps/boot"
	"github.com/auth/src/apps/boot/routes"
	"github.com/auth/src/apps/config"
	"github.com/auth/src/apps/controller"
	"github.com/auth/src/context/security"
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
		controller.ApiModule(),
		security.SecurityModule(),

		fx.Invoke(
			routes.SetRoutes,
			boot.StartHttpServer,
		),
	)
}
