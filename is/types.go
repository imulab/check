package is

import (
	"context"
	"github.com/imulab/check"
)

// String checks if the target is a string. If not, ErrType is returned.
var String check.Step = func(_ context.Context, target interface{}) error {
	_, ok := target.(string)
	if ok {
		return nil
	}

	return ErrType
}

// Int64 checks if the target is an int64. If not, ErrType is returned.
var Int64 check.Step = func(_ context.Context, target interface{}) error {
	_, ok := target.(int64)
	if ok {
		return nil
	}

	return ErrType
}

// Bool checks if the target is a bool. If not, ErrType is returned.
var Bool check.Step = func(_ context.Context, target interface{}) error {
	_, ok := target.(bool)
	if ok {
		return nil
	}

	return ErrType
}
