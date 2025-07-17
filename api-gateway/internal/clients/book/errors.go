package book

import "errors"

var (
	ErrBookNotFound       = errors.New("book not found")
	ErrInternalGRPCServer = errors.New("internal grpc server error")
)
