package usecase

import "context"

type SignUpConfirmUseCase interface {
	Execute(ctx context.Context, input SignUpConfirmInput) error
}

type signUpConfirmUseCase struct{}

func (s *signUpConfirmUseCase) Execute(ctx context.Context, input SignUpConfirmInput) error {
	return nil
}

type SignUpConfirmInput struct {
	Email            string
	ConfirmationCode string
}

func NewSignUpConfirmUseCase() SignUpConfirmUseCase {
	return &signUpConfirmUseCase{}
}
