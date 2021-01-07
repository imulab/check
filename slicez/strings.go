package slicez

import (
	"github.com/imulab/check"
	"github.com/imulab/check/stringz"
)

// OfString is the entry point for all check.Step which assumes the target is a string slice.
var OfString = stringTyped{
	IsEmpty: check.Step(func(target interface{}) error {
		if len(target.([]string)) == 0 {
			return nil
		}
		return ErrIsNotEmpty
	}),
	IsNotEmpty: check.Step(func(target interface{}) error {
		if len(target.([]string)) > 0 {
			return nil
		}
		return ErrIsEmpty
	}),
}

type stringTyped struct {
	// IsEmpty is a check.Step that verifies the target string slice
	// is empty, or returns ErrIsNotEmpty.
	IsEmpty check.Step
	// IsNotEmpty is a check.Step that verifies the target string slice
	// is not empty, or returns ErrIsEmpty.
	IsNotEmpty check.Step
}

// HasLength returns check.Step that verifies the given string slice has the expected length, or returns ErrHasLength.
func (_ stringTyped) HasLength(length int) check.Step {
	return func(target interface{}) error {
		if len(target.([]string)) == length {
			return nil
		}
		return ErrHasLength
	}
}

// HasLengthInRange returns check.Step that verifies the given string slice has the length in the
// expected range, or returns HasLengthInRange.
func (_ stringTyped) HasLengthInRange(startInclusive int, endExclusive int) check.Step {
	return func(target interface{}) error {
		length := len(target.([]string))
		if startInclusive <= length && length < endExclusive {
			return nil
		}
		return ErrHasLengthInRange
	}
}

// Contains returns check.Step that verifies the target string slice contains the expected element, or returns ErrContains.
func (s stringTyped) Contains(value string) check.Step {
	return s.Any(stringz.Is(value))
}

// NotContains returns check.Step that verifies the target string slice does not contain the element, or returns ErrNotContains.
func (s stringTyped) NotContains(value string) check.Step {
	return s.None(stringz.Is(value))
}

// All checks all string slice elements conform to the condition of the element check.Step. If an element check.Step
// returns an error, it is returned as the error. The element check.Step is NOT recommended to use check.Skip.
func (_ stringTyped) All(elemStep check.Step) check.Step {
	return func(target interface{}) error {
		for _, it := range target.([]string) {
			if err := elemStep(it); err != nil {
				return err
			}
		}
		return nil
	}
}

// Any checks if any string slice elements conform to the condition of the element check.Step. If all element
// check.Step returned error, ErrAny is returned.
func (_ stringTyped) Any(elemStep check.Step) check.Step {
	return func(target interface{}) error {
		for _, it := range target.([]string) {
			if err := elemStep(it); err == nil {
				return nil
			}
		}
		return ErrAny
	}
}

// None checks if none string slice elements conform to the condition of the element check.Step. If any element
// check.Step returned nil, ErrNone is returned.
func (_ stringTyped) None(elemStep check.Step) check.Step {
	return func(target interface{}) error {
		for _, it := range target.([]string) {
			if err := elemStep(it); err == nil {
				return ErrNone
			}
		}
		return nil
	}
}
