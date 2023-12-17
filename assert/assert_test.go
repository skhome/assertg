package assert

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

func assertErrorMessage(t *testing.T, fixture *fixtureT, message string) {
	t.Helper()
	if !strings.Contains(fixture.message, fmt.Sprintf("Error: %s", message)) {
		t.Errorf("expected to fail with error message %q, but got %#v", message, fixture.message)
	}
}

func assertErrorDescription(t *testing.T, fixture *fixtureT, description string) {
	t.Helper()
	if len(description) > 0 && !strings.Contains(fixture.message, fmt.Sprintf("Description: %s", description)) {
		t.Errorf("expected to fail with error description %q, but got %#v", description, fixture.message)
	}
}

func assertNoError(t *testing.T, fixture *fixtureT) {
	t.Helper()
	if len(fixture.message) > 0 {
		t.Errorf("expected not to fail, but got %#v", fixture.message)
	}
}
