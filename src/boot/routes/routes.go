package routes

import "github.com/auth/src/security/auth/api/auth_controller"

func SetRoutes(authRoutes *auth_controller.AuthRoutes) {
	authRoutes.Setup()
}
