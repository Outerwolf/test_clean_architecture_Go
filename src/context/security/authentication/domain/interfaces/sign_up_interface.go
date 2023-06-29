package interfaces

type ISignUpUseCase[T any] interface {
	Execute(request T) error
}
