package routes

import (
	"github.com/auth/src/apps/controller/authentication"
)

func SetRoutes(authRoutes *authentication.AuthRoutes) {
	authRoutes.Setup()
}
