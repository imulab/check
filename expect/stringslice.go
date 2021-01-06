package expect

import (
	"context"
	"errors"
	"github.com/imulab/check"
)

var (
	ErrMemberCondition     = errors.New("string slice member does not meet condition")
	ErrStringSliceEmpty    = errors.New("string slice is empty")
	ErrStringSliceNotEmpty = errors.New("string slice not empty")
	ErrStringSliceLength   = errors.New("string slice does not have expected length")
)

var StringSlice = stringSliceExpect{}

type stringSliceExpect struct{}

func (_ stringSliceExpect) ToBeOptional(_ context.Context, target interface{}) error {
	if len(target.([]string)) == 0 {
		return check.Skip
	}
	return nil
}

func (_ stringSliceExpect) ToBeEmpty(_ context.Context, target interface{}) error {
	if len(target.([]string)) == 0 {
		return nil
	}
	return ErrStringSliceNotEmpty
}

func (_ stringSliceExpect) ToNotBeEmpty(_ context.Context, target interface{}) error {
	if len(target.([]string)) > 0 {
		return nil
	}
	return ErrStringSliceEmpty
}

func (_ stringSliceExpect) ToHaveLength(length int) check.Step {
	return func(ctx context.Context, target interface{}) error {
		if len(target.([]string)) == length {
			return nil
		}
		return ErrStringSliceLength
	}
}

func (_ stringSliceExpect) ToHaveLengthInRange(startInclusive int, endExclusive int) check.Step {
	return func(ctx context.Context, target interface{}) error {
		l := len(target.([]string))
		if startInclusive <= l && l < endExclusive {
			return nil
		}
		return ErrStringSliceLength
	}
}

func (_ stringSliceExpect) All(memberStep check.Step) check.Step {
	return func(ctx context.Context, target interface{}) error {
		for _, it := range target.([]string) {
			err := memberStep(ctx, it)
			switch err {
			case nil:
				continue
			case check.Skip:
				return check.Skip
			default:
				return err
			}
		}
		return nil
	}
}

func (_ stringSliceExpect) Any(memberStep check.Step) check.Step {
	return func(ctx context.Context, target interface{}) error {
		for _, it := range target.([]string) {
			err := memberStep(ctx, it)
			switch err {
			case nil:
				return nil
			case check.Skip:
				return check.Skip
			default:
				continue
			}
		}
		return ErrMemberCondition
	}
}
