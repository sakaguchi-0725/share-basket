package usecase

import (
	"context"
)

type LoginUseCase interface {
	Execute(ctx context.Context, email, password string) (string, error)
}

type loginUseCase struct{}

func (l *loginUseCase) Execute(ctx context.Context, email string, password string) (string, error) {
	return "dummy-token", nil
}

func NewLoginUseCase() LoginUseCase {
	return &loginUseCase{}
}
