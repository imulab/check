package check_test

import (
	"context"
	"github.com/imulab/check"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpectType_ToBeString(t *testing.T) {
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
	}, check.Type.ToBeString)
}

func TestExpectType_ToBeInt64(t *testing.T) {
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
	}, check.Type.ToBeInt64)
}

func testType(t *testing.T, cases []typeTest, step check.Step) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := check.That(context.Background(), c.target, step)()
			if c.expect {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

type typeTest struct {
	name   string
	target interface{}
	expect bool
}
