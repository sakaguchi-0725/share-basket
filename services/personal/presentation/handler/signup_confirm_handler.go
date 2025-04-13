package handler

import (
	"context"
	"share-basket/personal/core"
	pb "share-basket/personal/gen/auth"
	"share-basket/personal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SignUpConfirmHandler interface {
	Handle(ctx context.Context, req *pb.SignUpConfirmRequest) (*emptypb.Empty, error)
}

type signUpConfirmHandler struct {
	usecase usecase.SignUpConfirmUseCase
}

func (h *signUpConfirmHandler) Handle(ctx context.Context, req *pb.SignUpConfirmRequest) (*emptypb.Empty, error) {
	err := h.usecase.Execute(ctx, usecase.SignUpConfirmInput{
		Email:            req.GetEmail(),
		ConfirmationCode: req.GetConfirmationCode(),
	})

	if err != nil {
		return nil, core.NewInvalidError(err)
	}

	return &emptypb.Empty{}, nil
}

func NewSignUpConfirmHandler(usecase usecase.SignUpConfirmUseCase) SignUpConfirmHandler {
	return &signUpConfirmHandler{usecase}
}
