package expect

import (
	"context"
	"errors"
	"github.com/imulab/check"
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
var String = stringExpect{
	ToBeOptional: check.Step(func(ctx context.Context, target interface{}) error {
		if len(target.(string)) == 0 {
			return check.Skip
		}
		return nil
	}),
	ToBeEmpty: check.Step(func(ctx context.Context, target interface{}) error {
		if len(target.(string)) == 0 {
			return nil
		}
		return ErrStringEmpty
	}),
	ToBeNonEmpty: check.Step(func(ctx context.Context, target interface{}) error {
		if len(target.(string)) == 0 {
			return nil
		}
		return ErrStringNotEmpty
	}),
}

type stringExpect struct {
	ToBeOptional check.Step
	ToBeEmpty    check.Step
	ToBeNonEmpty check.Step
}

func (_ stringExpect) ToEqual(value string) check.Step {
	return func(_ context.Context, target interface{}) error {
		if target.(string) == value {
			return nil
		}
		return ErrStringValue
	}
}

func (_ stringExpect) ToHaveValueIn(values ...string) check.Step {
	return func(_ context.Context, target interface{}) error {
		for _, it := range values {
			if it == target.(string) {
				return nil
			}
		}
		return ErrStringValue
	}
}

func (_ stringExpect) ToHaveLengthBetween(startInclusive int, endExclusive int) check.Step {
	return func(ctx context.Context, target interface{}) error {
		l := len(target.(string))
		if startInclusive <= l && l < endExclusive {
			return nil
		}
		return ErrStringLength
	}
}

func (_ stringExpect) ToMatch(regex regexp.Regexp) check.Step {
	return func(ctx context.Context, target interface{}) error {
		if regex.MatchString(target.(string)) {
			return nil
		}
		return ErrStringValue
	}
}
