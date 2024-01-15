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

func (f *fixtureT) Helper() {}

type DummyAssert struct {
	*BaseAssert[DummyAssert]
}

func TestDescription(t *testing.T) {
	fixture := new(fixtureT)
	baseAssert := &BaseAssert[DummyAssert]{t: fixture, info: NewWritableAssertionInfo()}

	baseAssert.DescribedAs("hobbit")
	baseAssert.FailWithMessage("message")

	if !strings.Contains(fixture.message, "[hobbit]") {
		t.Errorf("expected assertion error message to contain description %s, but got %s", "[hobbit]", fixture.message)
	}
}

func TestOverridingFailMessage(t *testing.T) {
	fixture := new(fixtureT)
	baseAssert := &BaseAssert[DummyAssert]{t: fixture, info: NewWritableAssertionInfo()}

	baseAssert.WithFailMessage("overriding message")
	baseAssert.FailWithMessage("default message")

	if !strings.Contains(fixture.message, "overriding message") {
		t.Errorf("expected assertion error message to contain %s, but got %s", "overriding message", fixture.message)
	}
}

func TestOverridingFailMessageSupplier(t *testing.T) {
	fixture := new(fixtureT)
	baseAssert := &BaseAssert[DummyAssert]{t: fixture, info: NewWritableAssertionInfo()}

	baseAssert.WithFailMessageSupplier(func() string { return "overriding message" })
	baseAssert.FailWithMessage("default message")

	if !strings.Contains(fixture.message, "overriding message") {
		t.Errorf("expected assertion error message to contain %s, but got %s", "overriding message", fixture.message)
	}
}

func TestRepresentation(t *testing.T) {
	fixture := new(fixtureT)
	baseAssert := &BaseAssert[DummyAssert]{t: fixture, info: NewWritableAssertionInfo()}

	currencyRepresentation := func(value any) string { return fmt.Sprintf("€%d", value) }
	baseAssert.WithRepresentation(currencyRepresentation)
	baseAssert.FailWithMessage("expected %s", 42)

	if !strings.Contains(fixture.message, "€42") {
		t.Errorf("expected assertion error message to contain %s, but got %s", "€42", fixture.message)
	}
}

func TestHexadecimalRepresentation(t *testing.T) {
	fixture := new(fixtureT)
	baseAssert := &BaseAssert[DummyAssert]{t: fixture, info: NewWritableAssertionInfo()}

	baseAssert.InHexadecimal()
	baseAssert.FailWithMessage("expected %s", 42)

	if !strings.Contains(fixture.message, "2A") {
		t.Errorf("expected assertion error message to contain %s, but got %s", "2A", fixture.message)
	}
}

func TestBinaryRepresentation(t *testing.T) {
	fixture := new(fixtureT)
	baseAssert := &BaseAssert[DummyAssert]{t: fixture, info: NewWritableAssertionInfo()}

	baseAssert.InBinary()
	baseAssert.FailWithMessage("expected %s", 42)

	if !strings.Contains(fixture.message, "101010") {
		t.Errorf("expected assertion error message to contain %s, but got %s", "101010", fixture.message)
	}
}
