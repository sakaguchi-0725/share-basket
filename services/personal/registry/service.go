package registry

import (
	"context"
	pb "share-basket/personal/gen/auth"
	"share-basket/personal/presentation/server"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewServices(c *container) server.Services {
	return server.Services{
		AuthService: newAuthService(c),
	}
}

type authService struct {
	container *container
	pb.UnimplementedAuthServiceServer
}

func (s *authService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.container.LoginHandler().Handle(ctx, req)
}

func (s *authService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*emptypb.Empty, error) {
	return s.container.SignUpHandler().Handle(ctx, req)
}

func (s *authService) SignUpConfirm(ctx context.Context, req *pb.SignUpConfirmRequest) (*emptypb.Empty, error) {
	return s.container.SignUpConfirmHandler().Handle(ctx, req)
}

func newAuthService(c *container) *authService {
	return &authService{container: c}
}
