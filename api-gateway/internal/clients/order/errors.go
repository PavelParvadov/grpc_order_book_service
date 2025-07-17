package order

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrBookIdNotFound      = errors.New("BookId Not Found")
)
