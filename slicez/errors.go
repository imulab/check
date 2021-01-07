package slicez

import "errors"

var (
	ErrIsEmpty          = errors.New("slice is not empty")
	ErrIsNotEmpty       = errors.New("slice is empty")
	ErrHasLength        = errors.New("slice does not have expected length")
	ErrHasLengthInRange = errors.New("slice does not have length in expected range")
	ErrContains         = errors.New("slice does not contain expected value")
	ErrNotContains      = errors.New("slice contains unexpected value")
	ErrAny              = errors.New("none of the slice elements meets to condition")
	ErrNone             = errors.New("some of the slice elements meets condition")
)
