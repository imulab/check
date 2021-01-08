package int64z_test

import (
	"github.com/imulab/check"
	"github.com/imulab/check/int64z"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		expect int64
		err    error
	}{
		{name: "1 equals 1", target: 1, expect: 1},
		{name: "1 does not equal 2", target: 1, expect: 2, err: int64z.ErrEquals},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.Equals(c.expect))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		expect int64
		err    error
	}{
		{name: "1 != 2", target: 1, expect: 2},
		{name: "1 == 1", target: 1, expect: 1, err: int64z.ErrNotEqual},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.NotEqual(c.expect))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}

func TestInRage(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		lower  int64
		upper  int64
		err    error
	}{
		{name: "in range", target: 3, lower: 1, upper: 10},
		{name: "=lower", target: 1, lower: 1, upper: 10},
		{name: "<lower", target: 0, lower: 1, upper: 10, err: int64z.ErrInRange},
		{name: "=upper", target: 10, lower: 1, upper: 10, err: int64z.ErrInRange},
		{name: ">lower", target: 11, lower: 1, upper: 10, err: int64z.ErrInRange},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.InRange(c.lower, c.upper))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		bound  int64
		err    error
	}{
		{name: "greater", target: 3, bound: 2},
		{name: "=bound", target: 2, bound: 2, err: int64z.ErrGreaterThan},
		{name: "<bound", target: 1, bound: 2, err: int64z.ErrGreaterThan},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.GreaterThan(c.bound))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}

func TestGreaterThanOrEqualTo(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		bound  int64
		err    error
	}{
		{name: "greater", target: 3, bound: 2},
		{name: "=bound", target: 2, bound: 2},
		{name: "<bound", target: 1, bound: 2, err: int64z.ErrGreaterThanOrEqualTo},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.GreaterThanOrEqualTo(c.bound))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		bound  int64
		err    error
	}{
		{name: "less", target: 1, bound: 2},
		{name: "=bound", target: 2, bound: 2, err: int64z.ErrLessThan},
		{name: ">bound", target: 3, bound: 2, err: int64z.ErrLessThan},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.LessThan(c.bound))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}

func TestLessThanOrEqualTo(t *testing.T) {
	cases := []struct {
		name   string
		target int64
		bound  int64
		err    error
	}{
		{name: "less", target: 1, bound: 2},
		{name: "=bound", target: 2, bound: 2},
		{name: ">bound", target: 3, bound: 2, err: int64z.ErrLessThanOrEqualTo},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(c.target, int64z.LessThanOrEqualTo(c.bound))()
			if c.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.err, err)
			}
		})
	}
}
