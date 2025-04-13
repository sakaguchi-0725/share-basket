package usecase

import "context"

type SignUpUseCase interface {
	Execute(ctx context.Context, input SignUpInput) error
}

type signUpUseCase struct{}

func (s *signUpUseCase) Execute(ctx context.Context, input SignUpInput) error {
	return nil
}

type SignUpInput struct {
	Name     string
	Email    string
	Password string
}

func NewSignUpUseCase() SignUpUseCase {
	return &signUpUseCase{}
}
