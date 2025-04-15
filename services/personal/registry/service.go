package registry

import (
	"context"
	"share-basket/personal/presentation/server"
	pb "share-basket/personal/proto/gen"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewServices(c *container) server.Services {
	return server.Services{
		AuthService:     newAuthService(c),
		ShoppingService: newShoppingService(c),
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

type shoppingService struct {
	container *container
	pb.UnimplementedPersonalShoppingServiceServer
}

func (s *shoppingService) Create(ctx context.Context, req *pb.CreateShoppingItemRequest) (*pb.ShoppingItem, error) {
	return s.container.CreateShoppingItemHandler().Handle(ctx, req)
}

func (s *shoppingService) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetShoppingItemsResponse, error) {
	return s.container.GetShoppingItemsHandler().Handle(ctx, req)
}

func newShoppingService(c *container) *shoppingService {
	return &shoppingService{container: c}
}
