package check

import (
	"errors"
)

var (
	// Skip is the special error to return in order to
	// skip the rest of validation.
	Skip = errors.New("skip")
)

// Step is a single validation step.
//
// Error returned will abort the remaining Step, and fail the validation run.
// A special case is Skip, return Skip is returned, the remaining Step is skipped and
// the validation run is treated as successful. Implementations are recommended to
// return a single documented error.
type Step func(target interface{}) error

// Err creates a new Step which returns the supplied error in case of failure.
func (s Step) Err(err error) Step {
	return func(target interface{}) error {
		se := s(target)
		switch se {
		case nil:
			return nil
		case Skip:
			return Skip
		default:
			return err
		}
	}
}

// If creates a new Step which executes this Step only when the condition Step returns nil.
func (s Step) If(obj interface{}, condition Step) Step {
	return func(target interface{}) error {
		ce := condition(obj)
		switch ce {
		case nil:
			return s(target)
		default:
			return nil
		}
	}
}

// When is like If, but applies the condition on the target itself.
func (s Step) When(condition Step) Step {
	return func(target interface{}) error {
		return s.If(target, condition)(target)
	}
}

// Optional is a Step that unconditionally emits Skip signal. It is useful to be
// combined with Step.If, or Step.When.
var Optional Step = func(_ interface{}) error {
	return Skip
}

// ErrFunc is a function that can return an error.
type ErrFunc func() error

// Err returns a wrapper ErrFunc to replace any returned error with the given error.
func (f ErrFunc) Err(err error) ErrFunc {
	return func() error {
		fe := f()
		if fe == nil {
			return nil
		}
		return err
	}
}

// That is the entrypoint for performing the validation Step. All supplied validation Step are
// performed sequentially unless an error is returned, or a Step returned Skip. That is a Validation.
func That(target interface{}, steps ...Step) ErrFunc {
	return func() error {
		for _, s := range steps {
			if err := s(target); err != nil {
				switch err {
				case Skip:
					return nil
				default:
					return err
				}
			}
		}
		return nil
	}
}

// AnyErr is a convenient invoker to chain multiple ErrFunc returned by That together.
func AnyErr(ef ...ErrFunc) error {
	for _, it := range ef {
		if err := it(); err != nil {
			return err
		}
	}
	return nil
}
