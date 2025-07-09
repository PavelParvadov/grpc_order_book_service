package service

import "errors"

var (
	ErrBookExists   = errors.New("book already exists")
	ErrBookNotFound = errors.New("book not found")
)
