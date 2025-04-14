package server

import (
	auth "share-basket/personal/gen/auth"
	personal "share-basket/personal/gen/personal"
)

type Services struct {
	AuthService     auth.AuthServiceServer
	ShoppingService personal.PersonalShoppingServiceServer
}
