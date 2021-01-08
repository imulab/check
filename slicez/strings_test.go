package slicez_test

import (
	"github.com/imulab/check"
	"github.com/imulab/check/slicez"
	"github.com/imulab/check/stringz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringTyped_HasLength(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		length int
		err    error
	}{
		{name: "has length", target: []string{"1", "2"}, length: 2},
		{name: "does not have length", target: []string{"1", "2"}, length: 3, err: slicez.ErrHasLength},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.HasLength(c.length))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringTyped_HasLengthInRange(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		lower  int
		upper  int
		err    error
	}{
		{name: "in range", target: []string{"1", "2"}, lower: 1, upper: 5},
		{name: "= lower", target: []string{"1", "2"}, lower: 2, upper: 5},
		{name: "< lower", target: []string{"1", "2"}, lower: 3, upper: 5, err: slicez.ErrHasLengthInRange},
		{name: "= upper", target: []string{"1", "2"}, lower: 1, upper: 2, err: slicez.ErrHasLengthInRange},
		{name: "> upper", target: []string{"1", "2"}, lower: 0, upper: 1, err: slicez.ErrHasLengthInRange},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.HasLengthInRange(c.lower, c.upper))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringTyped_Contains(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		seek   string
		err    error
	}{
		{name: "contains", target: []string{"1", "2"}, seek: "2"},
		{name: "does not contain", target: []string{"1", "2"}, seek: "3", err: slicez.ErrContains},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.Contains(c.seek))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringTyped_NotContain(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		seek   string
		err    error
	}{
		{name: "not contain", target: []string{"1", "2"}, seek: "3"},
		{name: "contains", target: []string{"1", "2"}, seek: "2", err: slicez.ErrNotContain},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.NotContain(c.seek))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringTyped_All(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		elem   check.Step
		err    error
	}{
		{name: "all", target: []string{"1", "2"}, elem: stringz.HasLength(1)},
		{name: "not all", target: []string{"1", "20"}, elem: stringz.HasLength(1), err: stringz.ErrHasLength},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.All(c.elem))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringTyped_Any(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		elem   check.Step
		err    error
	}{
		{name: "any", target: []string{"1", "20"}, elem: stringz.HasLength(1)},
		{name: "none", target: []string{"10", "20"}, elem: stringz.HasLength(1), err: slicez.ErrAny},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.Any(c.elem))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStringTyped_None(t *testing.T) {
	cases := []struct {
		name   string
		target []string
		elem   check.Step
		err    error
	}{
		{name: "none", target: []string{"10", "20"}, elem: stringz.HasLength(1)},
		{name: "one", target: []string{"1", "20"}, elem: stringz.HasLength(1), err: slicez.ErrNone},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, slicez.OfString.None(c.elem))()
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
