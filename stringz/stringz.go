package stringz

import (
	"errors"
	"github.com/imulab/check"
	"regexp"
	"strings"
)

var (
	ErrIs               = errors.New("unexpected string error")
	ErrIsEmpty          = errors.New("string is not empty")
	ErrIsNotEmpty       = errors.New("string is empty")
	ErrIn               = errors.New("string value not among expected values")
	ErrHasLength        = errors.New("string does not have expected length")
	ErrHasLengthInRange = errors.New("string does not have length in expected range")
	ErrHasPrefix        = errors.New("string does not have prefix")
	ErrHasSuffix        = errors.New("string does not have suffix")
	ErrContains         = errors.New("string does not contain expected value")
	ErrMatches          = errors.New("string does not match expected pattern")
)

// Is returns check.Step to verify target string has the expected value, or return ErrIs.
func Is(expect string) check.Step {
	return func(target interface{}) error {
		if target.(string) == expect {
			return nil
		}
		return ErrIs
	}
}

// IsEmpty is a check.Step that verifies the given string is empty (zero length), or returns ErrIsEmpty.
var IsEmpty check.Step = func(target interface{}) error {
	if len(target.(string)) == 0 {
		return nil
	}
	return ErrIsEmpty
}

// IsNotEmpty is a check.Step that verifies the given string is not empty, or returns ErrIsNotEmpty.
var IsNotEmpty check.Step = func(target interface{}) error {
	if len(target.(string)) > 0 {
		return nil
	}
	return ErrIsNotEmpty
}

// In returns a check.Step that verifies the target string value is among the expected list of values,
// or returns ErrIn.
func In(values ...string) check.Step {
	return func(target interface{}) error {
		for _, it := range values {
			if it == target.(string) {
				return nil
			}
		}
		return ErrIn
	}
}

// HasLength returns check.Step that verifies the given string has the expected length, or returns ErrHasLength.
func HasLength(length int) check.Step {
	return func(target interface{}) error {
		if len(target.(string)) == length {
			return nil
		}
		return ErrHasLength
	}
}

// HasLengthInRange returns check.Step that verifies the given string has the length in the
// expected range, or returns HasLengthInRange.
func HasLengthInRange(startInclusive int, endExclusive int) check.Step {
	return func(target interface{}) error {
		length := len(target.(string))
		if startInclusive <= length && length < endExclusive {
			return nil
		}
		return ErrHasLengthInRange
	}
}

// HasPrefix returns check.Step that verifies the target string has the expected prefix, or returns ErrHasPrefix.
func HasPrefix(prefix string) check.Step {
	return func(target interface{}) error {
		if strings.HasPrefix(target.(string), prefix) {
			return nil
		}
		return ErrHasPrefix
	}
}

// HasSuffix returns check.Step that verifies the target string has the expected suffix, or returns ErrHasSuffix.
func HasSuffix(suffix string) check.Step {
	return func(target interface{}) error {
		if strings.HasSuffix(target.(string), suffix) {
			return nil
		}
		return ErrHasSuffix
	}
}

// Contains returns check.Step that verifies the target string contains the expected substring, or returns ErrContains.
func Contains(substring string) check.Step {
	return func(target interface{}) error {
		if strings.Contains(target.(string), substring) {
			return nil
		}
		return ErrContains
	}
}

// Matches returns check.Step that verifies the target string matches the given regular expression pattern, or
// returns ErrMatches.
func Matches(pattern *regexp.Regexp) check.Step {
	return func(target interface{}) error {
		if pattern.MatchString(target.(string)) {
			return nil
		}
		return ErrMatches
	}
}
