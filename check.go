package check

import (
	"errors"
)

var (
	// Skip is the special error to return in order to skip the rest of validation.
	Skip = errors.New("skip")
)

// Step is a single validation step.
//
// Error returned will abort the remaining Step, and fail the validation run.
// A special case is Skip, return Skip is returned, the remaining Step is skipped and
// the validation run is treated as successful. Implementations are recommended to
// return a single documented error.
type Step func(target interface{}) error

// Err creates a new Step which returns the supplied error in case of failure. This is useful when uses want to
// supply a custom error as the validation error. Note that when Step returns Skip, it is not replaced.
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
//
//	// This example checks the one variable is "1", only if the two variable is "2".
//	func validate(one string, two string) error {
//		return check.That(one, stringz.Is("1").If(two, stringz.Is("2")))
//	}
//	validate("1", "2")	// returns nil (one is "1")
//	validate("x", "2")	// returns error (one is not "1")
//	validate("x", "y")	// returns nil (skipped because two is not "2")
//
// The condition Step is only accepted when it returns nil, any error (including Skip) will abort the dependent Step.
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
//
//	// This example states that str variable accepts either an empty string, or "foo".
//	check.That(str, stringz.Is("foo").When(stringz.IsNotEmpty))
func (s Step) When(condition Step) Step {
	return func(target interface{}) error {
		return s.If(target, condition)(target)
	}
}

// Optional is a Step that unconditionally emits Skip signal. It is useful to be
// combined with Step.If, or Step.When.
//
//	// This example skips checking str must be one of "a", "b" and "c", when it is empty.
//	check.That(str,
//		check.Optional.When(stringz.IsEmpty),
//		stringz.In("a", "b", "c"),
//	)
var Optional Step = func(_ interface{}) error {
	return Skip
}

// ErrFunc is a function that can return an error. ErrFunc works with AnyErr to provide
// a more fluent validation experience when involving multiple variables.
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
// performed sequentially unless an error is returned, or a Step returned Skip.
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
