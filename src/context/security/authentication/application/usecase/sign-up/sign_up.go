package sign_up

import (
	"fmt"
	"github.com/auth/src/apps/controller/authentication"
	"github.com/auth/src/context/security/authentication/domain/entity"
	"github.com/auth/src/context/security/authentication/domain/interfaces"
)

type SignUpUseCase struct {
}

func NewSignUpUseCase() interfaces.ISignUpUseCase[authentication.SignUpRequest] {
	return &SignUpUseCase{}
}

func (s *SignUpUseCase) Execute(request authentication.SignUpRequest) error {
	auth, err := entity.NewAuthentication(request.Id, request.UserName, request.Password)
	if err != nil {
		return err
	}
	fmt.Println(auth)
	return nil
}
