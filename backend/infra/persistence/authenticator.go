package persistence

import (
	"backend/domain/repository"
	"backend/infra/firebase"
	"context"

	"firebase.google.com/go/v4/errorutils"
)

type authenticator struct {
	firebaseClient firebase.Client
}

func (a *authenticator) VerifyToken(ctx context.Context, token string) (uid, email string, err error) {
	authToken, err := a.firebaseClient.VerifyIDToken(ctx, token)
	if err != nil {
		switch {
		case errorutils.IsInvalidArgument(err):
			return "", "", ErrInvalidToken
		case errorutils.IsUnauthenticated(err):
			return "", "", ErrTokenExpired
		case errorutils.IsPermissionDenied(err):
			return "", "", ErrAuthenticationFail
		default:
			return "", "", ErrAuthenticationFail
		}
	}

	uid = authToken.UID

	// emailが存在しない場合はエラー
	if emailValue, ok := authToken.Claims["email"].(string); !ok {
		return "", "", ErrInvalidToken
	} else {
		email = emailValue
	}

	return uid, email, nil
}

func NewAuthenticator(firebaseClient firebase.Client) repository.Authenticator {
	return &authenticator{firebaseClient: firebaseClient}
}
