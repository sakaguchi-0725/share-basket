package core

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidData = errors.New("invalid data")
	ErrNotFound    = errors.New("not found")
)

func NewInvalidError(err error) error {
	return status.Error(codes.InvalidArgument, err.Error())
}

func NewInternalError(err error) error {
	return status.Error(codes.Internal, err.Error())
}
