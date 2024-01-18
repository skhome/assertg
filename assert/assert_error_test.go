package assert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/skhome/assertg/assert"
)

type errorTest struct {
	err     error
	target  error
	message string
	substrs []string
	ok      bool
}

func TestErrorIsNil(t *testing.T) {
	tests := []errorTest{
		{err: nil, ok: true},
		{err: errors.New("error"), ok: false},
	}
	messageFormat := "expected error to be nil, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).IsNil()
		return test.ok, fmt.Sprintf(messageFormat, test.err)
	})
}

func TestErrorIsNotNil(t *testing.T) {
	tests := []errorTest{
		{err: errors.New("error"), ok: true},
		{err: nil, ok: false},
	}
	messageFormat := "expected error not to be nil, but got %v"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).IsNotNil()
		return test.ok, fmt.Sprintf(messageFormat, test.err)
	})
}

func TestErrorIs(t *testing.T) {
	cause := errors.New("cause")
	tests := []errorTest{
		{err: cause, target: cause, ok: true},
		{err: fmt.Errorf("failed because of: %w", cause), target: cause, ok: true},
		{err: errors.New("another"), target: cause, ok: false},
	}
	messageFormat := "expected error to have <%v> in its error chain, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).Is(test.target)
		return test.ok, fmt.Sprintf(messageFormat, test.target, test.err)
	})
}

func TestErrorIsNot(t *testing.T) {
	cause := errors.New("cause")
	tests := []errorTest{
		{err: errors.New("another"), target: cause, ok: true},
		{err: cause, target: cause, ok: false},
		{err: fmt.Errorf("failed because of: %w", cause), target: cause, ok: false},
	}
	messageFormat := "expected error not to have <%v> in its error chain, but got <%v>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).IsNot(test.target)
		return test.ok, fmt.Sprintf(messageFormat, test.target, test.err)
	})
}

func TestErrorHasMessage(t *testing.T) {
	err := errors.New("no such file")
	tests := []errorTest{
		{err: err, message: "no such file", ok: true},
		{err: err, message: "file", ok: false},
	}
	messageFormat := "expected error to have message <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).HasMessage(test.message)
		return test.ok, fmt.Sprintf(messageFormat, test.message, test.err)
	})
}

func TestErrorDoesNotHaveMessage(t *testing.T) {
	err := errors.New("no such file")
	tests := []errorTest{
		{err: err, message: "file", ok: true},
		{err: err, message: "no such file", ok: false},
	}
	messageFormat := "expected error not to have message <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).DoesNotHaveMessage(test.message)
		return test.ok, fmt.Sprintf(messageFormat, test.message, test.err)
	})
}

func TestErrorHasMessageContaining(t *testing.T) {
	err := errors.New("invalid users frodo, merry, pippin")
	tests := []errorTest{
		{err: err, substrs: []string{"frodo"}, ok: true},
		{err: err, substrs: []string{"frodo", "merry"}, ok: true},
		{err: err, substrs: []string{"sam"}, ok: false},
		{err: err, substrs: []string{"frodo", "sam"}, ok: false},
	}
	messageFormat := "expected error to have message containing <%v>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).HasMessageContaining(test.substrs...)
		return test.ok, fmt.Sprintf(messageFormat, test.substrs, test.err)
	})
}

func TestErrorHasMessageContainingAnyOf(t *testing.T) {
	err := errors.New("invalid users frodo, merry, pippin")
	tests := []errorTest{
		{err: err, substrs: []string{"frodo", "merry"}, ok: true},
		{err: err, substrs: []string{"frodo", "sam"}, ok: true},
		{err: err, substrs: []string{"sam"}, ok: false},
	}
	messageFormat := "expected error to have message containing any of <%v>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).HasMessageContainingAnyOf(test.substrs...)
		return test.ok, fmt.Sprintf(messageFormat, test.substrs, test.err)
	})
}

func TestErrorHasMessageNotContaining(t *testing.T) {
	err := errors.New("wrong amount 123")
	tests := []errorTest{
		{err: err, message: "234", ok: true},
		{err: err, message: "123", ok: false},
	}
	messageFormat := "expected error not to have message containing <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).HasMessageNotContaining(test.message)
		return test.ok, fmt.Sprintf(messageFormat, test.message, test.err)
	})
}

func TestErrorHasMessageStartingWith(t *testing.T) {
	err := errors.New("wrong amount 123")
	tests := []errorTest{
		{err: err, message: "wrong amount", ok: true},
		{err: err, message: "right amount", ok: false},
	}
	messageFormat := "expected error to have message starting with <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).HasMessageStartingWith(test.message)
		return test.ok, fmt.Sprintf(messageFormat, test.message, test.err)
	})
}

func TestErrorHasMessageEndingWith(t *testing.T) {
	err := errors.New("wrong amount 123")
	tests := []errorTest{
		{err: err, message: "123", ok: true},
		{err: err, message: "234", ok: false},
	}
	messageFormat := "expected error to have message ending with <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test errorTest) (bool, string) {
		assert.ThatError(fixture, test.err).HasMessageEndingWith(test.message)
		return test.ok, fmt.Sprintf(messageFormat, test.message, test.err)
	})
}
