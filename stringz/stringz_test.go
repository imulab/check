package stringz_test

import (
	"github.com/imulab/check"
	"github.com/imulab/check/stringz"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestIs(t *testing.T) {
	cases := []struct {
		name     string
		target   string
		expected string
		err      error
	}{
		{name: "foo == foo", target: "foo", expected: "foo"},
		{name: "foo != bar", target: "foo", expected: "bar", err: stringz.ErrIs},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.Is(c.expected))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		name   string
		target string
		err    error
	}{
		{name: "empty", target: ""},
		{name: "not empty", target: "foo", err: stringz.ErrIsEmpty},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.IsEmpty)()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		name   string
		target string
		err    error
	}{
		{name: "not empty", target: "foo"},
		{name: "empty", target: "", err: stringz.ErrIsNotEmpty},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.IsNotEmpty)()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIn(t *testing.T) {
	cases := []struct {
		name   string
		target string
		in     []string
		err    error
	}{
		{name: "in", target: "foo", in: []string{"baz", "foo", "bar"}},
		{name: "not in", target: "foo", in: []string{"baz", "bar"}, err: stringz.ErrIn},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.In(c.in...))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHasLength(t *testing.T) {
	cases := []struct {
		name   string
		target string
		length int
		err    error
	}{
		{name: "length equals", target: "foo", length: 3},
		{name: "length not equals", target: "foo", length: 4, err: stringz.ErrHasLength},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.HasLength(c.length))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHasLengthInRange(t *testing.T) {
	cases := []struct {
		name   string
		target string
		lower  int
		upper  int
		err    error
	}{
		{name: "length in range", target: "foo", lower: 1, upper: 5},
		{name: "length = lower", target: "foo", lower: 3, upper: 5},
		{name: "length < lower", target: "foo", lower: 4, upper: 5, err: stringz.ErrHasLengthInRange},
		{name: "length = upper", target: "foo", lower: 1, upper: 3, err: stringz.ErrHasLengthInRange},
		{name: "length > upper", target: "foo", lower: 1, upper: 2, err: stringz.ErrHasLengthInRange},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.HasLengthInRange(c.lower, c.upper))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	cases := []struct {
		name   string
		target string
		prefix string
		err    error
	}{
		{name: "is prefix", target: "foo", prefix: "f"},
		{name: "is not prefix", target: "foo", prefix: "o", err: stringz.ErrHasPrefix},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.HasPrefix(c.prefix))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	cases := []struct {
		name   string
		target string
		suffix string
		err    error
	}{
		{name: "is suffix", target: "foo", suffix: "o"},
		{name: "is not suffix", target: "foo", suffix: "f", err: stringz.ErrHasSuffix},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.HasSuffix(c.suffix))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name   string
		target string
		substr string
		err    error
	}{
		{name: "contains", target: "foo", substr: "o"},
		{name: "not contains", target: "foo", substr: "a", err: stringz.ErrContains},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.Contains(c.substr))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	cases := []struct {
		name    string
		target  string
		pattern string
		err     error
	}{
		{name: "matches", target: "foo", pattern: "^foo$"},
		{name: "does not match", target: "foo", pattern: "^bar$", err: stringz.ErrMatches},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, stringz.Matches(regexp.MustCompile(c.pattern)))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
