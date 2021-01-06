package check

import (
	"context"
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
type Step func(ctx context.Context, target interface{}) error

// Err creates a new Step which returns the supplied error in case of failure.
func (s Step) Err(err error) Step {
	return func(ctx context.Context, target interface{}) error {
		se := s(ctx, target)
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
func (s Step) If(condition Step) Step {
	return func(ctx context.Context, target interface{}) error {
		ce := condition(ctx, target)
		switch ce {
		case nil:
			return s(ctx, target)
		default:
			return nil
		}
	}
}

// ErrFunc is a function that can return an error.
type ErrFunc func() error

// That is the entrypoint for performing the validation Step. All supplied validation Step are
// performed sequentially unless an error is returned, or a Step returned Skip. That is a Validation.
func That(ctx context.Context, target interface{}, steps ...Step) ErrFunc {
	return func() error {
		for _, s := range steps {
			if err := s(ctx, target); err != nil {
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
