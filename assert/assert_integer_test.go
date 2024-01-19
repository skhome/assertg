package assert_test

import (
	"fmt"
	"testing"

	"github.com/skhome/assertg/assert"
	"golang.org/x/exp/constraints"
)

type integerTest[T constraints.Integer] struct {
	actual T
	other  T
	start  T
	end    T
	ok     bool
}

func TestIntegerIsEqualTo(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, other: 1, ok: true},
		{actual: 1, other: 2, ok: false},
	}
	messageFormat := "expected value to equal <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestIntegerIsNotEqualTo(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, other: -1, ok: true},
		{actual: 1, other: 1, ok: false},
	}
	messageFormat := "expected value not to equal <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsNotEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestIntegerZero(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 0, ok: true},
		{actual: 1, ok: false},
		{actual: -1, ok: false},
	}
	messageFormat := "expected value to be zero, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsZero()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerNonZero(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, ok: true},
		{actual: -1, ok: true},
		{actual: 0, ok: false},
	}
	messageFormat := "expected value not to be zero, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsNonZero()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerPositive(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, ok: true},
		{actual: 0, ok: false},
		{actual: -1, ok: false},
	}
	messageFormat := "expected value to be positive, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsPositive()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerNegative(t *testing.T) {
	tests := []integerTest[int]{
		{actual: -1, ok: true},
		{actual: 0, ok: false},
		{actual: 1, ok: false},
	}
	messageFormat := "expected value to be negative, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsNegative()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerIsNonPositive(t *testing.T) {
	tests := []integerTest[int]{
		{actual: -1, ok: true},
		{actual: 0, ok: true},
		{actual: 1, ok: false},
	}
	messageFormat := "expected value not to be positive, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsNonPostive()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerIsNonNegative(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, ok: true},
		{actual: 0, ok: true},
		{actual: -1, ok: false},
	}
	messageFormat := "expected value not to be negative, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsNonNegative()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerIsLessThan(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 0, other: 1, ok: true},
		{actual: 0, other: 0, ok: false},
	}
	messageFormat := "expected value to be less than <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsLessThan(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestIntegerIsLessThanOrEqualTo(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 0, other: 1, ok: true},
		{actual: 0, other: 0, ok: true},
		{actual: 1, other: 0, ok: false},
	}
	messageFormat := "expected value to be less than or equal to <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsLessThanOrEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestIntegerIsGreaterThan(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, other: 0, ok: true},
		{actual: 0, other: 0, ok: false},
	}
	messageFormat := "expected value to be greater than <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsGreaterThan(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestIntegerIsGreaterThanOrEqualTo(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, other: 1, ok: true},
		{actual: 1, other: 0, ok: true},
		{actual: 0, other: 1, ok: false},
	}
	messageFormat := "expected value to be greater than or equal to <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsGreaterThanOrEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.actual)
	})
}

func TestIntegerIsBetween(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, start: 0, end: 2, ok: true},
		{actual: 1, start: 1, end: 2, ok: true},
		{actual: 1, start: 0, end: 1, ok: true},
		{actual: 1, start: 2, end: 3, ok: false},
	}
	messageFormat := "expected value to be between <%d> and <%d>, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsBetween(test.start, test.end)
		return test.ok, fmt.Sprintf(messageFormat, test.start, test.end, test.actual)
	})
}

func TestIntegerIsEven(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 0, ok: true},
		{actual: 2, ok: true},
		{actual: 1, ok: false},
	}
	messageFormat := "expected value to be even, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsEven()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}

func TestIntegerIsOdd(t *testing.T) {
	tests := []integerTest[int]{
		{actual: 1, ok: true},
		{actual: 0, ok: false},
		{actual: 2, ok: false},
	}
	messageFormat := "expected value to be odd, but got <%d>"
	runTests(t, tests)(func(fixture *fixtureT, test integerTest[int]) (bool, string) {
		assert.ThatInteger(fixture, test.actual).IsOdd()
		return test.ok, fmt.Sprintf(messageFormat, test.actual)
	})
}
