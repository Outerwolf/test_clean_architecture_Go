package auth_controller

import "github.com/auth/src/boot"

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
	s.handler.PUT("/api/security/auth/sign-up/:id", s.authController.SignUp)
}
