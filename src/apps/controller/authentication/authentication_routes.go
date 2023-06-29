package authentication

import (
	"github.com/auth/src/apps/boot"
)

type AuthRoutes struct {
	handler        *boot.Router
	authController *AuthController
}

func NewAuthRoutes(handler *boot.Router, authController *AuthController) *AuthRoutes {
	return &AuthRoutes{
		handler:        handler,
		authController: authController,
	}
}

func (s *AuthRoutes) Setup() {
	s.handler.POST("/api/security/authentication/sign-up", s.authController.SignUp)
}
