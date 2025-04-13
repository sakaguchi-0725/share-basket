package handler

import (
	"context"
	"share-basket/personal/core"
	pb "share-basket/personal/gen/auth"
	"share-basket/personal/usecase"
)

type LoginHandler interface {
	Handle(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
}

type loginHandler struct {
	usecase usecase.LoginUseCase
}

func (l *loginHandler) Handle(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := l.usecase.Execute(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, core.NewInternalError(err)
	}

	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}

func NewLoginHandler(usecase usecase.LoginUseCase) LoginHandler {
	return &loginHandler{usecase}
}
