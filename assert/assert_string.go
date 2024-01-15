package assert

import (
	"regexp"
	"slices"
	"strings"

	check "github.com/skhome/assertg/check"
)

// StringAssert provides assertions on strings.
type StringAssert struct {
	*BaseAssert[StringAssert]
	actual string
}

// IsEmpty verifies that the actual string is empty, i.e. has a length of 0.
//
//	// assertion fill pass
//	 assert.ThatString(t, "").IsEmpty()
//
//	// assertion will fail
//	assert.ThatString(t, " ").IsEmpty()
//	assert.ThatString(t, "a").IsEmpty()
func (a *StringAssert) IsEmpty() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) > 0 {
		a.FailWithMessage("expected string to be empty, but got %s", a.actual)
	}
	return a
}

// IsNotEmpty verifies that the actual string is not empty, i.e. has a length of 1 or more.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").IsNotEmpty()
//	assert.ThatString(t, " ").IsNotEmpty()
//
//	// assertion will fail
//	assert.ThatString(t, "").IsNotEmpty()
func (a *StringAssert) IsNotEmpty() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) == 0 {
		a.FailWithMessage("expected string to not be empty, but got %s", a.actual)
	}
	return a
}

// IsBlank verifies that the actual string is blank, i.e. empty or contains only whitespace characters.
//
//	// assertion will pass
//	assert.ThatString(t, "").IsBlank()
//	assert.ThatString(t, " \t").IsBlank()
//
//	// assertion will fail
//	assert.ThatString(t, "a").IsBlank()
func (a *StringAssert) IsBlank() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringIsBlank(a.actual) {
		a.FailWithMessage("expected string to be blank, but got %s", a.actual)
	}
	return a
}

// IsNotBlank verifies that the actual string is not blank, i.e. not empty and contains at least one non-whitespace character.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").IsNotBlank()
//
//	// assertion will fail
//	assert.ThatString(t, "").IsNotBlank() // FAIL
//	assert.ThatString(t, " ").IsNotBlank() // FAIL
func (a *StringAssert) IsNotBlank() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringIsBlank(a.actual) {
		a.FailWithMessage("expected string to not be blank, but got %s", a.actual)
	}
	return a
}

// ContainsWhitespace verifies that the actual string contains one or more whitespace characters.
//
//	// assertion will pass
//	assert.ThatString(t, " ").ContainsWhitespace()
//	assert.ThatString(t, "Frodo Baggins").ContainsWhitespace()
//
//	// assertion will fail
//	assert.ThatString(t, "").ContainsWhitespace()
//	assert.ThatString(t, "Frodo").ContainsWhitespace()
func (a *StringAssert) ContainsWhitespace() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsWhitespace(a.actual) {
		a.FailWithMessage("expected string to contain whitespace characters, but got %s", a.actual)
	}
	return a
}

// DoesNotContainWhitespace verifies that the actual string does not contain any whitespace characters.
//
//	// assertion will pass
//	assert.ThatString(t, "").DoesNotContainWhitespace()
//	assert.ThatString(t, "Frodo").DoesNotContainWhitespace()
//
//	// assertion will fail
//	assert.ThatString(t, " ").DoesNotContainWhitespace()
//	assert.ThatString(t, "Frodo Baggins").DoesNotContainWhitespace()
func (a *StringAssert) DoesNotContainWhitespace() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringContainsWhitespace(a.actual) {
		a.FailWithMessage("expected string to not contain whitespace characters, but got %s", a.actual)
	}
	return a
}

// HasLength verifies that the actual string has the expected length.
//
//	// assertion will pass
//	assert.ThatString(t, "Lord of the Rings").HasLength(17)
//
//	// assertion will fail
//	assert.ThatString(t, "Lord of the RIngs").HasLength(10)
func (a *StringAssert) HasLength(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != length {
		a.FailWithMessage("expected string to have length of %s, but got %s", length, a.actual)
	}
	return a
}

// HasLengthLessThan verifies that the actual string has a length less than the given value.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").HasLengthLessThan(6)
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").HasLengthLessThan(5)
func (a *StringAssert) HasLengthLessThan(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) >= length {
		a.FailWithMessage("expected string to have length less than %s, but got %s", length, a.actual)
	}
	return a
}

// HasLengthGreaterThan verifies that the actual string has a length greater than the given value.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").HasLengthGreaterThan(4)
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").HasLengthGreaterThan(5)
func (a *StringAssert) HasLengthGreaterThan(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) <= length {
		a.FailWithMessage("expected string to have length greater than %s, but got %s", length, a.actual)
	}
	return a
}

// HasSameLengthAs verifies that the actual string has the same length as the given string.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf").HasSameLengthAs("Saruman")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf").HasSameLengthAs("Frodo")
func (a *StringAssert) HasSameLengthAs(other string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != len(other) {
		a.FailWithMessage("expected string to have the same length as %s, but got %s", other, a.actual)
	}
	return a
}

// HasLineCount verifies that the actual string has the expected line count.
//
//	// assertion will pass
//	assert.ThatString(t, "first\nsecond").HasLineCount(2)
//
//	// assertion will fail
//	assert.ThatString(t, "first\nsecond\n").HasLineCount(2)
func (a *StringAssert) HasLineCount(expectedLineCount int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringLineCount(a.actual) != expectedLineCount {
		a.FailWithMessage("expected string to have %s lines, but got %s", expectedLineCount, a.actual)
	}
	return a
}

// IsEqualTo verifies that the actual string equals the given one.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").IsEqualTo("Frodo")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").IsEqualTo("Sam")
func (a *StringAssert) IsEqualTo(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if expected != a.actual {
		a.FailWithMessage("expected string to equal %s, but got %s", expected, a.actual)
	}
	return a
}

// IsNotEqualTo verifies that the actual string does not equal the given one.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").IsNotEqualTo("Sam")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").IsNotEqualTo("Frodo")
func (a *StringAssert) IsNotEqualTo(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if expected == a.actual {
		a.FailWithMessage("expected string not to equal %s, but got %s", expected, a.actual)
	}
	return a
}

// IsEqualToIgnoringCase verifies that the actual string equals the given one, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").IsEqualToIgnoringCase("frodo")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").IsEqualToIgnoringCase("gandalf")
func (a *StringAssert) IsEqualToIgnoringCase(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.EqualFold(expected, a.actual) {
		a.FailWithMessage("expected string to equal %s ignoring case, but got %s", expected, a.actual)
	}
	return a
}

// IsNotEqualToIgnoringCase verifies that the actual string is not equal the given one, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf").IsNotEqualToIgnoringCase("Frodo")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf").IsNotEqualToIgnoringCase("gandalf")
func (a *StringAssert) IsNotEqualToIgnoringCase(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.EqualFold(expected, a.actual) {
		a.FailWithMessage("expected string to not equal %s ignoring case, but got %s", expected, a.actual)
	}
	return a
}

// ContainsDigit verifies that the actual string contains only digits.
//
//	// assertion will pass
//	assert.ThatString(t, "Bug8ear").ContainsDigit()
//
//	// assertion will fail
//	assert.ThatString(t, "V").ContainsDigit()
func (a *StringAssert) ContainsDigit() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsDigit(a.actual) {
		a.FailWithMessage("expected string to contain digit, but got %s", a.actual)
	}
	return a
}

// ContainsOnlyDigits verifies that the actual string contains only digits.
//
//	// assertion will pass
//	assert.ThatString(t, "10").
//	       ContainsOnlyDigits()
//
//	// assertion will fail
//	assert.ThatString(t, "10€").
//	       ContainsOnlyDigits()
func (a *StringAssert) ContainsOnlyDigits() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsOnlyDigits(a.actual) {
		a.FailWithMessage("expected string to only contain digits, but got %s", a.actual)
	}
	return a
}

// ContainsOnlyOnce verifies that the actual string contains the given substring only once.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       ContainsOnlyOnce("do")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       ContainsOnlyOnce("o").
//	       ContainsOnlyOnce("y")
func (a *StringAssert) ContainsOnlyOnce(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.Count(a.actual, substr) != 1 {
		a.FailWithMessage("expected string to contain %s only once, but got %s", substr, a.actual)
	}
	return a
}

// Contains verifies that the actual string contains all the given values as substring.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       Contains("alf", "grey")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       Contains("white")
func (a *StringAssert) Contains(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContains(a.actual, values) {
		a.FailWithMessage("expected string to contain %s, but got %s", values, a.actual)
	}
	return a
}

// ContainsAnyOf verifies that the actual string contains any of the given substrings.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       ContainsAnyOf("grey", "black")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       ContainsAnyOf("white", "black")
func (a *StringAssert) ContainsAnyOf(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsAny(a.actual, values) {
		a.FailWithMessage("expected string to contain any of %s, but got %s", values, a.actual)
	}
	return a
}

// ContainsIgnoringCase verifies that the actual string contains the given substring, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       ContainsIgnoringCase("gandalf", "Grey")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       ContainsIgnoringCase("white")
func (a *StringAssert) ContainsIgnoringCase(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsIgnoringCase(a.actual, values) {
		a.FailWithMessage("expected string to contain %s ignoring case, but got %s", values, a.actual)
	}
	return a
}

// ContainsIgnoringWhitespace verifies that the actual string contains the given substring, ignoring whitespace characters.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       ContainsIgnoringWhitespace("alf").
//	       ContainsIgnoringWhitespace("alf", "grey").
//	       ContainsIgnoringWhitespace("thegrey").
//	       ContainsIgnoringWhitespace("thegr  ey").
//	       ContainsIgnoringWhitespace("t h e\t grey")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       ContainsIgnoringWhitespace("alF")
func (a *StringAssert) ContainsIgnoringWhitespace(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringContainsIgnoringWhitespace(a.actual, values) {
		a.FailWithMessage("expected string to contain %s ignoring whitespace, but got %s", values, a.actual)
	}
	return a
}

// DoesNotContain verifies that the actual string does not contain the given substring.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotContain("pippin").
//	       DoesNotContain("fro", "sam")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotContain("Fro", "Gimli")
func (a *StringAssert) DoesNotContain(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringContainsAny(a.actual, values) {
		a.FailWithMessage("expected string to not contain %s, but got %s", values, a.actual)
	}
	return a
}

// DoesNotContainIgnoringCase verifies that the actual string does not contain the given substring, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotContainIgnoringCase("pippin").
//	       DoesNotContainIgnoringCase("Merry", "sam")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotContainIgnoringCase("Fro", "Gimli").
//	       DoesNotContainIgnoringCase("fro")
func (a *StringAssert) DoesNotContainIgnoringCase(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringContainsAnyIgnoringCase(a.actual, values) {
		a.FailWithMessage("expected string to not contain %s ignoring case, but got %s", values, a.actual)
	}
	return a
}

// DoesNotContainIgnoringWhitespace verifies that the actual string does not contain the given substring, ignoring whitespace.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       DoesNotContainIgnoringWhitespace("TheGrey")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       DoesNotContainIgnoringWhitespace("thegrey")
func (a *StringAssert) DoesNotContainIgnoringWhitespace(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringContainsAnyIgnoringWhitespace(a.actual, values) {
		a.FailWithMessage("expected string to not contain %s ignoring whitespace, but got %s", values, a.actual)
	}
	return a
}

// StartsWith verifies that the actual string starts with the given prefix.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       StartsWith("Fro")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       StartsWith("fro")
func (a *StringAssert) StartsWith(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringStartsWith(a.actual, prefix) {
		a.FailWithMessage("expected string to start with %s, but got %s", prefix, a.actual)
	}
	return a
}

// DoesNotStartWith verifies that the actual string does not start with the given prefix.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotStartWith("fro")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotStartWith("")
func (a *StringAssert) DoesNotStartWith(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringStartsWith(a.actual, prefix) {
		a.FailWithMessage("expected string not to start with %s, but got %s", prefix, a.actual)
	}
	return a
}

// StartsWithIgnoringCase verifies that the actual string starts with the given prefix, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       StartsWithIgnoringCase("Gandalf").
//	       StartsWithIgnoringCase("gandalf")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       StartsWith("grey")
func (a *StringAssert) StartsWithIgnoringCase(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringStartsWithIgnoringCase(a.actual, prefix) {
		a.FailWithMessage("expected string to start with %s ignoring case, but got %s", prefix, a.actual)
	}
	return a
}

// DoesNotStartWithIgnoringCase verifies that the actual string does not start with the given prefix, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Gandalf the grey").
//	       DoesNotStartWithIgnoringCase("fro").
//	       DoesNotStartWithIgnoringCase("grey")
//
//	// assertion will fail
//	assert.ThatString(t, "Gandalf the grey").
//	       DoesNotStartWithIgnoringCase("Gandalf").
//	       DoesNotStartWithIgnoringCase("gandalf")
func (a *StringAssert) DoesNotStartWithIgnoringCase(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringStartsWithIgnoringCase(a.actual, prefix) {
		a.FailWithMessage("expected string not to start with %s ignoring case, but got %s", prefix, a.actual)
	}
	return a
}

// EndsWith verifies that the actual string ends with the given suffix.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       EndsWith("do")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       EndsWith("Fro")
func (a *StringAssert) EndsWith(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringEndsWith(a.actual, suffix) {
		a.FailWithMessage("expected string to end with %s, but got %s", suffix, a.actual)
	}
	return a
}

// DoesNotEndWith verifies that the actual string does not end with the given suffix.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotEndWith("Fro")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotEndWith("do").
//	       DoesNotEndWith("")
func (a *StringAssert) DoesNotEndWith(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringEndsWith(a.actual, suffix) {
		a.FailWithMessage("expected string not to end with %s, but got %s", suffix, a.actual)
	}
	return a
}

// EndsWithIgnoringCase verifies that the actual string ends with the given suffix, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       EndsWithIgnoringCase("do").
//	       EndsWithIgnoringCase("Do")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       EndsWithIgnoringCase("fro")
func (a *StringAssert) EndsWithIgnoringCase(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringEndsWithIgnoringCase(a.actual, suffix) {
		a.FailWithMessage("expected string to end with %s ignoring case, but got %s", suffix, a.actual)
	}
	return a
}

// DoesNotEndWithIgnoringCase verifies that the actual string does not end with the given suffix, ignoring case considerations.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotEndWithIgnoringCase("Fro")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotEndWithIgnoringCase("do").
//	       DoesNotEndWithIgnoringCase("DO")
func (a *StringAssert) DoesNotEndWithIgnoringCase(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringEndsWithIgnoringCase(a.actual, suffix) {
		a.FailWithMessage("expected string not to end with %s ignoring case, but got %s", suffix, a.actual)
	}
	return a
}

// MatchesPattern verifies that the actual string matches the given regular expression pattern.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       MatchesPattern(`..o.o`)
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       MatchesPattern(`.*d$`)
func (a *StringAssert) MatchesPattern(pattern string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringMatchesRegexp(a.actual, regexp.MustCompile(pattern)) {
		a.FailWithMessage("expected string to match %s, but got %s", pattern, a.actual)
	}
	return a
}

// DoesNotMatchPattern verifies that the actual string does not match the given regular expression.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotMatchPattern(`.*d$`)
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotMatchPattern(`..o.o`)
func (a *StringAssert) DoesNotMatchPattern(pattern string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringMatchesRegexp(a.actual, regexp.MustCompile(pattern)) {
		a.FailWithMessage("expected string not to match %s, but got %s", pattern, a.actual)
	}
	return a
}

// MatchesRegexp verifies that the actual string matches the given compiled regular expression.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       MatchesRegexp(regexp.MustCompile(`..o.o`))
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       MatchesRegexp(regexp.MustCompile(`.*f`))
func (a *StringAssert) MatchesRegexp(re *regexp.Regexp) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringMatchesRegexp(a.actual, re) {
		a.FailWithMessage("expected string to match %s, but got %s", re, a.actual)
	}
	return a
}

// DoesNotMatchRegexp verifies that the actual string does not match the given compiled regular expression.
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       DoesNotMatchRegexp(regexp.MustCompile(`.*d$`))
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       DoesNotMatchRegexp(regexp.MustCompile(`..o.o`))
func (a *StringAssert) DoesNotMatchRegexp(re *regexp.Regexp) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringMatchesRegexp(a.actual, re) {
		a.FailWithMessage("expected string to match %s, but got %s", re, a.actual)
	}
	return a
}

// IsEqualToIgnoringWhitespace verifies that the actual string is equal to the given one, ignoring whitespace differences.
//
//	// assertion will pass
//	assert.ThatString(t, "Game of Thrones").
//	       IsEqualToIgnoringWhitespace("Game   of   Thrones").
//	       IsEqualToIgnoringWhitespace("  Game of   Thrones  ").
//	       IsEqualToIgnoringWhitespace("  Game of Thrones  ").
//	       IsEqualToIgnoringWhitespace("\tGame of Thrones\n")
//
//	// assertion will fail
//	assert.ThatString(t, "Game of Thrones").
//	       IsEqualToIgnoringWhitespace("Game OF Thrones")
func (a *StringAssert) IsEqualToIgnoringWhitespace(value string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !check.StringEqualsIgnoringWhitespace(a.actual, value) {
		a.FailWithMessage("expected string to equal ignoring whitespace %s, but got %s", value, a.actual)
	}
	return a
}

// IsNotEqualToIgnoringWhitespace verifies that the actual string does not equal the given one, ignoring whitespace differences.
//
//	// assertion will pass
//	assert.ThatString(t, "Game of Thrones").
//	       IsNotEqualToIgnoringWhitespace("Game OF Thrones")
//
//	// assertion will fail
//	assert.ThatString(t, "Game of Thrones").
//	       IsNotEqualToIgnoringWhitespace("Game   of   Thrones").
//	       IsNotEqualToIgnoringWhitespace("  Game of   Thrones  ").
//	       IsNotEqualToIgnoringWhitespace("  Game of Thrones  ").
//	       IsNotEqualToIgnoringWhitespace("\tGame of Thrones\n")
func (a *StringAssert) IsNotEqualToIgnoringWhitespace(value string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if check.StringEqualsIgnoringWhitespace(a.actual, value) {
		a.FailWithMessage("expected string not to equal ignoring whitespace %s, but got %s", value, a.actual)
	}
	return a
}

// IsSubstringOf verifies that the actual string is a substring of the given string.
//
//	// assertion will pass
//	assert.ThatString(t, "Lego").
//	       IsSubstringOf("Legolas").
//	       IsSubstringOf("Lego")
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       IsSubstringOf("Frod")
func (a *StringAssert) IsSubstringOf(str string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.Contains(str, a.actual) {
		a.FailWithMessage("expected string to be a substring of %s, but got %s", str, a.actual)
	}
	return a
}

// IsIn verifies that the actual string is present in the given slice.
//
//	hobbits := []string{"Frodo", "Sam", "Merry", "Pippin", "Bilbo"}
//
//	// assertion will pass
//	assert.ThatString(t, "Frodo").
//	       IsIn(hobbits)
//
//	// assertion will fail
//	assert.ThatString(t, "Legolas").
//	       IsIn(hobbits)
func (a *StringAssert) IsIn(slice []string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !slices.Contains(slice, a.actual) {
		a.FailWithMessage("expected string to be present in %s, but got %s", slice, a.actual)
	}
	return a
}

// IsNotIn verifies that the actual string is not present in the given slice.
//
//	hobbits := []string{"Frodo", "Sam", "Merry", "Pippin", "Bilbo"}
//
//	// assertion will pass
//	assert.ThatString(t, "Legolas").
//	       IsIn(hobbits)
//
//	// assertion will fail
//	assert.ThatString(t, "Frodo").
//	       IsIn(hobbits)
func (a *StringAssert) IsNotIn(slice []string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if slices.Contains(slice, a.actual) {
		a.FailWithMessage("expected string not to be present in %s, but got %s", slice, a.actual)
	}
	return a
}

// IsLowerCase verifies that is actual string is all lower case.
//
//	// assertions will pass
//	assert.ThatString(t, "legolas").IsLowerCase()
//	assert.ThatString(t, "").IsLowerCase()
//	assert.ThatString(t, ".").IsLowerCase()
//	assert.ThatString(t, "42").IsLowerCase()
//
//	// assertions will fail
//	assert.ThatString(t, "Legolas").IsLowerCase()
func (a *StringAssert) IsLowerCase() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual != strings.ToLower(a.actual) {
		a.FailWithMessage("expected string to be all lower case, but got %s", a.actual)
	}
	return a
}

// IsUpperCase verifies that is actual string is all upper case.
//
//	// assertions will pass
//	assert.ThatString(t, "LEGOLAS").IsUpperCase()
//	assert.ThatString(t, "").IsUpperCase()
//	assert.ThatString(t, ".").IsUpperCase()
//	assert.ThatString(t, "42").IsUpperCase()
//
//	// assertions will fail
//	assert.ThatString(t, "Legolas").IsUpperCase()
func (a *StringAssert) IsUpperCase() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if a.actual != strings.ToUpper(a.actual) {
		a.FailWithMessage("expected string to be all upper case, but got %s", a.actual)
	}
	return a
}
