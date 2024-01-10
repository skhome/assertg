package assert_test

import (
	"fmt"
	"strings"
	"testing"
)

type fixtureT struct {
	message string
}

func (f *fixtureT) Errorf(format string, args ...any) {
	f.message = fmt.Sprintf(format, args...)
}

func (f *fixtureT) Helper() {}

// function that runs a single test
type testExecutor[E any] func(fixture *fixtureT, test E) (bool, string)

// run a list of tests and verify error messages
func runTests[E any](t *testing.T, tests []E) func(execTest testExecutor[E]) {
	t.Helper()
	return func(execTest testExecutor[E]) {
		t.Helper()
		for _, test := range tests {
			fixture := new(fixtureT)
			if ok, message := execTest(fixture, test); ok {
				assertNoError(t, fixture)
			} else {
				assertErrorMessage(t, fixture, message)
			}
		}
	}
}

func assertErrorMessage(t *testing.T, fixture *fixtureT, message string) {
	t.Helper()
	if !strings.Contains(fixture.message, fmt.Sprintf("Error: %s", message)) {
		t.Errorf("expected to fail with error message %q, but got %#v", message, fixture.message)
	}
}

func assertNoError(t *testing.T, fixture *fixtureT) {
	t.Helper()
	if len(fixture.message) > 0 {
		t.Errorf("expected not to fail, but got %#v", fixture.message)
	}
}
