package expect

import (
	"context"
	"errors"
)

var (
	ErrTypeNotString = errors.New("target is not string typed")
	ErrTypeNotInt64  = errors.New("target is not int64 typed")
)

// Type is the entrypoint for checking types.
var Type = expectType{}

type expectType struct{}

func (_ expectType) ToBeString(_ context.Context, target interface{}) error {
	if _, ok := target.(string); ok {
		return nil
	}
	return ErrTypeNotString
}

func (_ expectType) ToBeInt64(_ context.Context, target interface{}) error {
	if _, ok := target.(int64); ok {
		return nil
	}
	return ErrTypeNotInt64
}
