package sign_up

import (
	"github.com/auth/src/security/auth/api/auth_controller"
	"github.com/auth/src/security/auth/domain/interfaces"
)

type SignUpUseCase struct {
}

func NewSignUpUseCase() interfaces.ISignUpUseCase[auth_controller.SignUpRequest] {
	return &SignUpUseCase{}
}

func (s *SignUpUseCase) Execute(request auth_controller.SignUpRequest) {
	return
}
