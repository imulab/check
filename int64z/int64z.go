package int64z

import (
	"errors"
	"github.com/imulab/check"
)

var (
	ErrEquals               = errors.New("int64 value does not equal to expected value")
	ErrNotEqual             = errors.New("int64 value equals unexpected value")
	ErrInRange              = errors.New("int64 value is not in range")
	ErrGreaterThan          = errors.New("int64 value is not greater than expected value")
	ErrLessThan             = errors.New("int64 value is not ")
	ErrGreaterThanOrEqualTo = errors.New("int64 value is less than expected value")
	ErrLessThanOrEqualTo    = errors.New("int64 value is greater than expected value")
)

var (
	// Zero is a convenient check.Step to check equality to 0
	Zero = Equals(0)
	// Positive is a convenient check.Step to check greater than 0
	Positive = GreaterThan(0)
	// Negative is a convenient check.Step to check less than 0
	Negative = LessThan(0)
	// NonPositive is a convenient check.Step to check less than or equal to 0
	NonPositive = LessThanOrEqualTo(0)
	// NonNegative is a convenient check.Step to check greater than or equal to 0
	NonNegative = GreaterThanOrEqualTo(0)
)

// Equals returns a check.Step that check equality of target and expected value, or returns ErrEquals.
func Equals(expected int64) check.Step {
	return func(target interface{}) error {
		if expected == target.(int64) {
			return nil
		}
		return ErrEquals
	}
}

// NotEqual returns a check.Step to check inequality to target to the value, or returns ErrNotEqual.
func NotEqual(unexpected int64) check.Step {
	return func(target interface{}) error {
		if unexpected != target.(int64) {
			return nil
		}
		return ErrNotEqual
	}
}

// InRange returns a check.Step that check the int64 target value is in the given range, specified by an
// inclusive start value and an exclusive end value, or returns ErrInRange.
func InRange(startInclusive int64, endExclusive int64) check.Step {
	return func(target interface{}) error {
		i := target.(int64)
		if startInclusive <= i && i < endExclusive {
			return nil
		}
		return ErrInRange
	}
}

// GreaterThan returns a check.Step that check the int64 target value is greater than the expected bound value,
// or returns a ErrGreaterThan
func GreaterThan(bound int64) check.Step {
	return func(target interface{}) error {
		if target.(int64) > bound {
			return nil
		}
		return ErrGreaterThan
	}
}

// LessThan returns a check.Step that check the int64 target value is less than the expected bound value,
// or returns a ErrLessThan
func LessThan(bound int64) check.Step {
	return func(target interface{}) error {
		if target.(int64) < bound {
			return nil
		}
		return ErrLessThan
	}
}

// GreaterThanOrEqualTo returns a check.Step that check the int64 target value is greater than or equal to
// the expected bound value, or returns a ErrGreaterThanOrEqualTo
func GreaterThanOrEqualTo(bound int64) check.Step {
	return func(target interface{}) error {
		if target.(int64) >= bound {
			return nil
		}
		return ErrGreaterThanOrEqualTo
	}
}

// LessThanOrEqualTo returns a check.Step that check the int64 target value is less than or equal to
// the expected bound value, or returns a LessThanOrEqualTo.
func LessThanOrEqualTo(bound int64) check.Step {
	return func(target interface{}) error {
		if target.(int64) <= bound {
			return nil
		}
		return ErrLessThanOrEqualTo
	}
}
