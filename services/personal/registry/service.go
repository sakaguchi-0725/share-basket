package registry

import (
	"context"
	auth "share-basket/personal/gen/auth"
	personal "share-basket/personal/gen/personal"
	"share-basket/personal/presentation/server"

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
	auth.UnimplementedAuthServiceServer
}

func (s *authService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return s.container.LoginHandler().Handle(ctx, req)
}

func (s *authService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*emptypb.Empty, error) {
	return s.container.SignUpHandler().Handle(ctx, req)
}

func (s *authService) SignUpConfirm(ctx context.Context, req *auth.SignUpConfirmRequest) (*emptypb.Empty, error) {
	return s.container.SignUpConfirmHandler().Handle(ctx, req)
}

func newAuthService(c *container) *authService {
	return &authService{container: c}
}

type shoppingService struct {
	container *container
	personal.UnimplementedPersonalShoppingServiceServer
}

func (s *shoppingService) Create(ctx context.Context, req *personal.CreateShoppingItemRequest) (*personal.ShoppingItem, error) {
	return s.container.CreateShoppingItemHandler().Handle(ctx, req)
}

func (s *shoppingService) GetAll(ctx context.Context, req *emptypb.Empty) (*personal.GetShoppingItemsResponse, error) {
	return s.container.GetShoppingItemsHandler().Handle(ctx, req)
}

func newShoppingService(c *container) *shoppingService {
	return &shoppingService{container: c}
}
