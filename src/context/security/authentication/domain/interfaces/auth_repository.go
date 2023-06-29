package interfaces

import (
	"github.com/auth/src/context/security/authentication/domain/entity"
)

type IAuthenticationRepository interface {
	getById(id string) entity.Authentication
	create(request entity.Authentication)
	update(id string, request entity.Authentication)
	delete(id string)
}
