package assert_test

import (
	"fmt"
	"testing"

	"github.com/skhome/assertg/assert"
)

type boolTest struct {
	value    bool
	expected bool
	ok       bool
}

func TestBoolIsTrue(t *testing.T) {
	tests := []boolTest{
		{value: true, ok: true},
		{value: false, ok: false},
	}
	messageFormat := "expected value to be true, but got <%t>"
	runTests(t, tests)(func(fixture *fixtureT, test boolTest) (bool, string) {
		assert.ThatBool(fixture, test.value).IsTrue()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestBoolIsFalse(t *testing.T) {
	tests := []boolTest{
		{value: false, ok: true},
		{value: true, ok: false},
	}
	messageFormat := "expected value to be false, but got <%t>"
	runTests(t, tests)(func(fixture *fixtureT, test boolTest) (bool, string) {
		assert.ThatBool(fixture, test.value).IsFalse()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestBoolIsEqualTo(t *testing.T) {
	tests := []boolTest{
		{value: false, expected: false, ok: true},
		{value: true, expected: false, ok: false},
	}
	messageFormat := "expected value to be <%t>, but got <%t>"
	runTests(t, tests)(func(fixture *fixtureT, test boolTest) (bool, string) {
		assert.ThatBool(fixture, test.value).IsEqualTo(test.expected)
		return test.ok, fmt.Sprintf(messageFormat, test.expected, test.value)
	})
}

func TestBoolIsNotEqualTo(t *testing.T) {
	tests := []boolTest{
		{value: false, expected: true, ok: true},
		{value: true, expected: true, ok: false},
	}
	messageFormat := "expected value not to be <%t>, but got <%t>"
	runTests(t, tests)(func(fixture *fixtureT, test boolTest) (bool, string) {
		assert.ThatBool(fixture, test.value).IsNotEqualTo(test.expected)
		return test.ok, fmt.Sprintf(messageFormat, test.expected, test.value)
	})
}
