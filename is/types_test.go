package is_test

import (
	"context"
	"github.com/imulab/check"
	"github.com/imulab/check/is"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsString(t *testing.T) {
	testType(t, []typeTest{
		{
			name:   "string",
			target: "foo",
			expect: true,
		},
		{
			name:   "not a string",
			target: 123,
			expect: false,
		},
	}, is.String)
}

func TestIsInt64(t *testing.T) {
	testType(t, []typeTest{
		{
			name:   "int64",
			target: int64(123),
			expect: true,
		},
		{
			name:   "not an int64",
			target: "foo",
			expect: false,
		},
	}, is.Int64)
}

func TestIsBool(t *testing.T) {
	testType(t, []typeTest{
		{
			name:   "bool",
			target: true,
			expect: true,
		},
		{
			name:   "not a bool",
			target: "foo",
			expect: false,
		},
	}, is.Bool)
}

func testType(t *testing.T, cases []typeTest, step check.Step) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(context.Background(), c.target, step)
			if c.expect {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, is.ErrType, err)
			}
		})
	}
}

type typeTest struct {
	name   string
	target interface{}
	expect bool
}
