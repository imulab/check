package expect

import (
	"context"
	"errors"
	"github.com/imulab/check"
)

var (
	ErrInt64Value = errors.New("int64 value is invalid")
)

var Int64 = int64Expect{
	ToBePositive: check.Step(func(ctx context.Context, target interface{}) error {
		if target.(int64) > 0 {
			return nil
		}
		return ErrInt64Value
	}),
	ToBeNonNegative: check.Step(func(ctx context.Context, target interface{}) error {
		if target.(int64) >= 0 {
			return nil
		}
		return ErrInt64Value
	}),
}

type int64Expect struct {
	ToBePositive    check.Step
	ToBeNonNegative check.Step
}

func (_ int64Expect) ToEqual(i int64) check.Step {
	return func(ctx context.Context, target interface{}) error {
		if target.(int64) == i {
			return nil
		}
		return ErrInt64Value
	}
}

func (_ int64Expect) ToNotEqual(i int64) check.Step {
	return func(ctx context.Context, target interface{}) error {
		if target.(int64) != i {
			return nil
		}
		return ErrInt64Value
	}
}

func (_ int64Expect) ToBeInRange(startInclusive int64, endExclusive int64) check.Step {
	return func(ctx context.Context, target interface{}) error {
		i := target.(int64)
		if startInclusive <= i && i < endExclusive {
			return nil
		}
		return ErrInt64Value
	}
}
