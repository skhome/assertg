package assert_test

import (
	"fmt"
	"testing"

	"github.com/skhome/assertg/assert"
	"golang.org/x/exp/constraints"
)

type floatTest[T constraints.Float] struct {
	actual T
	other  T
	start  T
	end    T
	ok     bool
}

func TestFloatIsEqualTo(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, other: 1, ok: true},
		{actual: 1.25, other: 1.25, ok: true},
		{actual: 1, other: 2, ok: false},
	}
	messageFormat := "expected value to equal <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestFloatIsNotEqualTo(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, other: -1, ok: true},
		{actual: 1, other: 1, ok: false},
	}
	messageFormat := "expected value not to equal <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsNotEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestFloatZero(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 0, ok: true},
		{actual: 1, ok: false},
		{actual: -1, ok: false},
	}
	messageFormat := "expected value to be zero, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsZero()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestFloatNonZero(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, ok: true},
		{actual: -1, ok: true},
		{actual: 0, ok: false},
	}
	messageFormat := "expected value not to be zero, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsNonZero()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestFloatPositive(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, ok: true},
		{actual: 0, ok: false},
		{actual: -1, ok: false},
	}
	messageFormat := "expected value to be positive, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsPositive()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestFloatNegative(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: -1, ok: true},
		{actual: 0, ok: false},
		{actual: 1, ok: false},
	}
	messageFormat := "expected value to be negative, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsNegative()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestFloatIsNonPositive(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: -1, ok: true},
		{actual: 0, ok: true},
		{actual: 1, ok: false},
	}
	messageFormat := "expected value not to be positive, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsNonPostive()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestFloatIsNonNegative(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, ok: true},
		{actual: 0, ok: true},
		{actual: -1, ok: false},
	}
	messageFormat := "expected value not to be negative, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsNonNegative()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestFloatIsLessThan(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 0, other: 1, ok: true},
		{actual: 0, other: 0, ok: false},
	}
	messageFormat := "expected value to be less than <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsLessThan(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestFloatIsLessThanOrEqualTo(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 0, other: 1, ok: true},
		{actual: 0, other: 0, ok: true},
		{actual: 1, other: 0, ok: false},
	}
	messageFormat := "expected value to be less than or equal to <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsLessThanOrEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestFloatIsGreaterThan(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, other: 0, ok: true},
		{actual: 0, other: 0, ok: false},
	}
	messageFormat := "expected value to be greater than <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsGreaterThan(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestFloatIsGreaterThanOrEqualTo(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, other: 1, ok: true},
		{actual: 1, other: 0, ok: true},
		{actual: 0, other: 1, ok: false},
	}
	messageFormat := "expected value to be greater than or equal to <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsGreaterThanOrEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestFloatIsBetween(t *testing.T) {
	tests := []floatTest[float32]{
		{actual: 1, start: 0, end: 2, ok: true},
		{actual: 1, start: 1, end: 2, ok: true},
		{actual: 1, start: 0, end: 1, ok: true},
		{actual: 1, start: 2, end: 3, ok: false},
	}
	messageFormat := "expected value to be between <%v> and <%v>, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test floatTest[float32]) (bool, string) {
		assert.ThatFloat(fixture, test.actual).IsBetween(test.start, test.end)
		return test.ok, fmt.Sprintf(messageFormat, test.start, test.end, test.actual)
	})
}
