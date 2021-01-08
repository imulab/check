package check_test

import (
	"errors"
	"github.com/imulab/check"
	"github.com/imulab/check/stringz"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	correctStep check.Step = func(target interface{}) error {
		return nil
	}
	wrongStep check.Step = func(target interface{}) error {
		return errors.New("step has error")
	}
)

func TestStep_If(t *testing.T) {
	assert.Error(t, wrongStep.If("foo", stringz.HasLength(3))("anything"))
	assert.NoError(t, wrongStep.If("foo", stringz.HasLength(5))("anything"))
}

func TestStep_When(t *testing.T) {
	assert.Error(t, wrongStep.When(stringz.HasLength(3))("foo"))
	assert.NoError(t, wrongStep.When(stringz.HasLength(5))("foo"))
}

func TestStep_Err(t *testing.T) {
	var customErr = errors.New("customErr")
	assert.Equal(t, customErr, wrongStep.Err(customErr)("anything"))
}

func TestErrFunc_Err(t *testing.T) {
	var customErr = errors.New("customErr")
	assert.Equal(t, customErr, check.That("foo", wrongStep).Err(customErr)())
}

func TestOptional(t *testing.T) {
	err := check.AnyErr(
		check.That("foo", check.Optional.When(correctStep), wrongStep),
	)
	assert.NoError(t, err)
}

func TestAnyErr(t *testing.T) {
	err := check.AnyErr(
		check.That("foo", correctStep),
		check.That("foo", wrongStep),
	)
	assert.Error(t, err)
}
