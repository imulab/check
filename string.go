package check

import (
	"context"
	"errors"
	"regexp"
)

var (
	ErrStringValue    = errors.New("string value is invalid")
	ErrStringEmpty    = errors.New("string value is empty")
	ErrStringNotEmpty = errors.New("string value is not empty")
	ErrStringLength   = errors.New("string length is out of bounds")
)

// String is the entrypoint for checking strings. All check.Step
// under here expects the target to be string. If caller is unsure
// of the type, call Type.ToBeString first.
var String = stringExpect{}

type stringExpect struct{}

func (_ stringExpect) ToHaveValueIn(values ...string) Step {
	return func(_ context.Context, target interface{}) error {
		for _, it := range values {
			if it == target.(string) {
				return nil
			}
		}
		return ErrStringValue
	}
}

func (_ stringExpect) ToBeEmpty(_ context.Context, target interface{}) error {
	if len(target.(string)) == 0 {
		return nil
	}
	return ErrStringEmpty
}

func (_ stringExpect) ToBeNonEmpty(_ context.Context, target interface{}) error {
	if len(target.(string)) == 0 {
		return nil
	}
	return ErrStringNotEmpty
}

func (_ stringExpect) ToHaveLengthBetween(include int, exclude int) Step {
	return func(ctx context.Context, target interface{}) error {
		l := len(target.(string))
		if include <= l && l < exclude {
			return nil
		}
		return ErrStringLength
	}
}

func (_ stringExpect) ToMatch(regex regexp.Regexp) Step {
	return func(ctx context.Context, target interface{}) error {
		if regex.MatchString(target.(string)) {
			return nil
		}
		return ErrStringValue
	}
}
