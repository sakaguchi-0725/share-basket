package registry

import (
	"share-basket/personal/presentation/handler"
	"share-basket/personal/usecase"
)

type container struct {
	loginUseCase         usecase.LoginUseCase
	signUpUseCase        usecase.SignUpUseCase
	signUpConfirmUseCase usecase.SignUpConfirmUseCase

	createShoppingItemUseCase usecase.CreateShoppingItemUseCase
	getShoppingItemsUseCase   usecase.GetShoppingItemsUseCase
}

func Inject() (*container, error) {
	return &container{
		loginUseCase:              usecase.NewLoginUseCase(),
		signUpUseCase:             usecase.NewSignUpUseCase(),
		signUpConfirmUseCase:      usecase.NewSignUpConfirmUseCase(),
		createShoppingItemUseCase: usecase.NewCreateShoppingItemUseCase(),
		getShoppingItemsUseCase:   usecase.NewGetShoppingItemsUseCase(),
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

func (c *container) CreateShoppingItemHandler() handler.CreateShoppingItemHandler {
	return handler.NewCreateShoppingItemHandler(c.createShoppingItemUseCase)
}

func (c *container) GetShoppingItemsHandler() handler.GetShoppingItemsHandler {
	return handler.NewGetShoppingItemsHandler(c.getShoppingItemsUseCase)
}
