package server

import (
	pb "share-basket/personal/proto/gen"
)

type Services struct {
	AuthService     pb.AuthServiceServer
	ShoppingService pb.PersonalShoppingServiceServer
}
