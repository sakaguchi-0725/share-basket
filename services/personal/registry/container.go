package registry

import (
	"share-basket/personal/presentation/handler"
	"share-basket/personal/usecase"
)

type container struct {
	loginUseCase         usecase.LoginUseCase
	signUpUseCase        usecase.SignUpUseCase
	signUpConfirmUseCase usecase.SignUpConfirmUseCase
}

func Inject() (*container, error) {
	return &container{
		loginUseCase:         usecase.NewLoginUseCase(),
		signUpUseCase:        usecase.NewSignUpUseCase(),
		signUpConfirmUseCase: usecase.NewSignUpConfirmUseCase(),
	}, nil
}

func (c *container) LoginHandler() handler.LoginHandler {
	return handler.NewLoginHandler(c.loginUseCase)
}

func (c *container) SignUpHandler() handler.SignUpHandler {
	return handler.NewSignUpHandler(c.signUpUseCase)
}

func (c *container) SignUpConfirmHandler() handler.SignUpConfirmHandler {
	return handler.NewSignUpConfirmHandler(c.signUpConfirmUseCase)
}
