package check

import (
	"context"
	"errors"
)

var (
	ErrMemberCondition = errors.New("string slice member does not meet condition")
)

var StringSlice = stringSliceExpect{}

type stringSliceExpect struct{}

// todo empty, not empty, length, lengthInRange

func (_ stringSliceExpect) All(memberStep Step) Step {
	return func(ctx context.Context, target interface{}) error {
		for _, it := range target.([]string) {
			err := memberStep(ctx, it)
			switch err {
			case nil:
				continue
			case Skip:
				return Skip
			default:
				return err
			}
		}
		return nil
	}
}

func (_ stringSliceExpect) Any(memberStep Step) Step {
	return func(ctx context.Context, target interface{}) error {
		for _, it := range target.([]string) {
			err := memberStep(ctx, it)
			switch err {
			case nil:
				return nil
			case Skip:
				return Skip
			default:
				continue
			}
		}
		return ErrMemberCondition
	}
}
