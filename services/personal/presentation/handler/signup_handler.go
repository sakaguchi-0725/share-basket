package handler

import (
	"context"
	"share-basket/personal/core"
	pb "share-basket/personal/gen/auth"
	"share-basket/personal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SignUpHandler interface {
	Handle(ctx context.Context, req *pb.SignUpRequest) (*emptypb.Empty, error)
}

type signUpHandler struct {
	usecase usecase.SignUpUseCase
}

func (h *signUpHandler) Handle(ctx context.Context, req *pb.SignUpRequest) (*emptypb.Empty, error) {
	input := h.makeInput(req)
	if err := h.usecase.Execute(ctx, input); err != nil {
		return nil, core.NewInvalidError(err)
	}

	return &emptypb.Empty{}, nil
}

func (h *signUpHandler) makeInput(req *pb.SignUpRequest) usecase.SignUpInput {
	return usecase.SignUpInput{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func NewSignUpHandler(usecase usecase.SignUpUseCase) SignUpHandler {
	return &signUpHandler{usecase}
}
