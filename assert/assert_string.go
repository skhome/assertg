package assert

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// StringAssert provides assertions on strings.
type StringAssert struct {
	t           TestingT
	message     string
	description string
	actual      string
}

// As sets an optional description for this assertion.
func (a *StringAssert) As(format string, args ...any) *StringAssert {
	a.description = fmt.Sprintf(format, args...)
	return a
}

// WithFailMessage overrides the failure message.
func (a *StringAssert) WithFailMessage(format string, args ...any) *StringAssert {
	a.message = fmt.Sprintf(format, args...)
	return a
}

// failWithMessage records an assertion failure.
func (a *StringAssert) failWithMessage(format string, args ...any) {
	message := a.message
	if len(message) == 0 {
		message = fmt.Sprintf(format, args...)
	}
	Fail(a.t, message, a.description)
}

// IsEmpty verifies that the actual string is empty, i.e. has a length of 0.
//
// Examples:
//
//	assert.ThatString(t, "").IsEmpty() // OK
//	assert.ThatString(t, " ").IsEmpty() // FAIL
//	assert.ThatString(t, "a").IsEmpty() // FAIL
func (a *StringAssert) IsEmpty() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) > 0 {
		a.failWithMessage("expected string to be empty, but got %q", a.actual)
	}
	return a
}

// IsNotEmpty verifies that the actual string is not empty, i.e. has a length of 1 or more.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").IsNotEmpty() // OK
//	assert.ThatString(t, " ").IsNotEmpty() // OK
//	assert.ThatString(t, "").IsNotEmpty() // FAIL
func (a *StringAssert) IsNotEmpty() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) == 0 {
		a.failWithMessage("expected string to not be empty, but got %q", a.actual)
	}
	return a
}

// IsBlank verifies that the actual string is blank, i.e. empty or contains only whitespace characters.
//
// Examples:
//
//	assert.ThatString(t, "").IsBlank() // OK
//	assert.ThatString(t, " \t").IsBlank() // OK
//	assert.ThatString(t, " Frodo").IsBlank() // FAIL
func (a *StringAssert) IsBlank() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(strings.TrimSpace(a.actual)) > 0 {
		a.failWithMessage("expected string to be blank, but got %q", a.actual)
	}
	return a
}

// IsNotBlank verifies that the actual string is not blank, i.e. not empty and contains at least one non-whitespace character.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").IsNotBlank() // OK
//	assert.ThatString(t, "").IsNotBlank() // FAIL
//	assert.ThatString(t, " ").IsNotBlank() // FAIL
func (a *StringAssert) IsNotBlank() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(strings.TrimSpace(a.actual)) == 0 {
		a.failWithMessage("expected string to not be blank, but got %q", a.actual)
	}
	return a
}

// ContainsWhitespace verifies that the actual string contains one or more whitespace characters.
//
// Examples:
//
//	assert.ThatString(t, " ").ContainsWhitespace() // OK
//	assert.ThatString(t, "Frodo Baggins").ContainsWhitespace() // OK
//	assert.ThatString(t, "").ContainsWhitespace() // FAIL
//	assert.ThatString(t, "Frodo").ContainsWhitespace() // FAIL
func (a *StringAssert) ContainsWhitespace() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.ContainsFunc(a.actual, unicode.IsSpace) {
		a.failWithMessage("expected string to contain whitespace characters, but got %q", a.actual)
	}
	return a
}

// DoesNotContainWhitespace verifies that the actual string does not contain any whitespace characters.
//
// Examples:
//
//	assert.ThatString(t, "").DoesNotContainWhitespace() // OK
//	assert.ThatString(t, "Frodo").DoesNotContainWhitespace() // OK
//	assert.ThatString(t, " ").DoesNotContainWhitespace() // FAIL
//	assert.ThatString(t, "Frodo Baggins").DoesNotContainWhitespace() // FAIL
func (a *StringAssert) DoesNotContainWhitespace() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.ContainsFunc(a.actual, unicode.IsSpace) {
		a.failWithMessage("expected string to not contain whitespace characters, but got %q", a.actual)
	}
	return a
}

// HasLength verifies that the actual string has the expected length.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").HasLength(5) // OK
//	assert.ThatString(t, "Frodo").HasLength(4) // FAIL
func (a *StringAssert) HasLength(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != length {
		a.failWithMessage("expected string to have length of %d, but got %q", length, a.actual)
	}
	return a
}

// HasLengthLessThan verifies that the actual string has a length less than expected.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").HasLengthLessThan(6) // OK
//	assert.ThatString(t, "Frodo").HasLengthLessThan(5) // FAIL
func (a *StringAssert) HasLengthLessThan(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) >= length {
		a.failWithMessage("expected string to have length less than %d, but got %q", length, a.actual)
	}
	return a
}

// HasLengthGreaterThan verifies that the actual string has a length greater than expected.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").HasLengthGreaterThan(4) // OK
//	assert.ThatString(t, "Frodo").HasLengthGreaterThan(5) // FAIL
func (a *StringAssert) HasLengthGreaterThan(length int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) <= length {
		a.failWithMessage("expected string to have length greater than %d, but got %q", length, a.actual)
	}
	return a
}

// HasLineCount verifies that the actual string has the expected line count.
//
// Examples:
//
//	assert.ThatString(t, "first\nsecond").HasLineCount(2) // OK
//	assert.ThatString(t, "first\nsecond\n").HasLengthGreaterThan(2) // FAIL
func (a *StringAssert) HasLineCount(expectedLineCount int) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	actualLineCount := strings.Count(a.actual, "\n") + 1
	if expectedLineCount != actualLineCount {
		a.failWithMessage("expected string to have %d lines, but got %q", expectedLineCount, a.actual)
	}
	return a
}

// HasSameLengthAs verifies that the actual string has the same length as the given string.
//
// Examples:
//
//	assert.ThatString(t, "Gandalf").HasSameLengthAs("Saruman") // OK
//	assert.ThatString(t, "Gandalf").HasSameLengthAs("Frodo") // FAIL
func (a *StringAssert) HasSameLengthAs(other string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if len(a.actual) != len(other) {
		a.failWithMessage("expected string to have the same length as %q (%d), but got %q (%d)", other, len(other), a.actual, len(a.actual))
	}
	return a
}

// IsEqualTo verifies that the actual string equals the given one.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").IsEqualTo("Frodo") // OK
//	assert.ThatString(t, "Frodo").IsEqualTo("Sam") // FAIL
func (a *StringAssert) IsEqualTo(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if expected != a.actual {
		a.failWithMessage("expected string to equal %q, but got %q", expected, a.actual)
	}
	return a
}

// IsNotEqualTo verifies that the actual string does not equal the given one.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").IsNotEqualTo("Sam") // OK
//	assert.ThatString(t, "Frodo").IsNotEqualTo("Frodo") // FAIL
func (a *StringAssert) IsNotEqualTo(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if expected == a.actual {
		a.failWithMessage("expected string to not equal %q, but got %q", expected, a.actual)
	}
	return a
}

// IsEqualToIgnoringCase verifies that the actual string equals the given one, ignoring case considerations.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").IsEqualToIgnoringCase("frodo") // OK
//	assert.ThatString(t, "frodo").IsEqualToIgnoringCase("gandalf") // FAIL
func (a *StringAssert) IsEqualToIgnoringCase(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.EqualFold(expected, a.actual) {
		a.failWithMessage("expected string to equal %q ignoring case, but got %q", expected, a.actual)
	}
	return a
}

// IsNotEqualToIgnoringCase verifies that the actual string is not equal the given one, ignoring case considerations.
//
// Examples:
//
//	assert.ThatString(t, "Gandalf").IsNotEqualToIgnoringCase("Hobbit") // OK
//	assert.ThatString(t, "Gandalf").IsNotEqualToIgnoringCase("gandalf") // FAIL
func (a *StringAssert) IsNotEqualToIgnoringCase(expected string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.EqualFold(expected, a.actual) {
		a.failWithMessage("expected string to not equal %q ignoring case, but got %q", expected, a.actual)
	}
	return a
}

// ContainsOnlyDigits verifies that the actual string contains only digits.
//
// Examples:
//
//	assert.ThatString(t, "10").ContainsOnlyDigits() // OK
//	assert.ThatString(t, "10€").ContainsOnlyDigits() // FAIL
func (a *StringAssert) ContainsOnlyDigits() *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	foundNonDigit := false
	for _, r := range a.actual {
		if !unicode.IsDigit(r) {
			foundNonDigit = true
			break
		}
	}
	if foundNonDigit {
		a.failWithMessage("expected string to only contain digits, but got %q", a.actual)
	}
	return a
}

// ContainsOnlyOnce verifies that the actual string contains the given substring only once.
//
// Examples:
//
//	assert.ThatString(t, "Frodo").ContainsOnlyOnce("do") // OK
//	assert.ThatString(t, "Frodo").ContainsOnlyOnce("o") // FAIL
//	assert.ThatString(t, "Frodo").ContainsOnlyOnce("y") // FAIL
func (a *StringAssert) ContainsOnlyOnce(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	firstIndex := strings.Index(a.actual, substr)
	lastIndex := strings.LastIndex(a.actual, substr)
	if firstIndex == -1 || firstIndex != lastIndex {
		a.failWithMessage("expected string to contain %q only once, but got %q", substr, a.actual)
	}
	return a
}

// Contains verifies that the actual string contains the given substring.
//
// Examples:
//
//	assert.ThatString(t, "Gandalf the grey").Contains("grey") // OK
//	assert.ThatString(t, "Gandalf the grey").Contains("white") // FAIL
func (a *StringAssert) Contains(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.Contains(a.actual, substr) {
		a.failWithMessage("expected string to contain %q, but got %q", substr, a.actual)
	}
	return a
}

// ContainsIgnoringCase verifies that the actual string contains the given substring, ignoring case considerations.
//
// Examples:
//
//	assert.ThatString(t, "Gandalf the grey").ContainsIgnoringCase("gandalf") // OK
//	assert.ThatString(t, "Gandalf the grey").ContainsIgnoringCase("white") // FAIL
func (a *StringAssert) ContainsIgnoringCase(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.Contains(strings.ToLower(a.actual), strings.ToLower(substr)) {
		a.failWithMessage("expected string to contain %q ignoring case, but got %q", substr, a.actual)
	}
	return a
}

func removeWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// ContainsIgnoringWhitespace verifies that the actual string contains the given substring, ignoring whitespace characters.
//
// Examples:
//
//	assert.ThatString(t, "Gandalf the grey").ContainsIgnoringWhitespace("thegrey") // OK
//	assert.ThatString(t, "Gandalf the grey").ContainsIgnoringWhitespace("Grey") // FAIL
func (a *StringAssert) ContainsIgnoringWhitespace(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.Contains(removeWhitespace(a.actual), removeWhitespace(substr)) {
		a.failWithMessage("expected string to contain %q ignoring whitespace, but got %q", substr, a.actual)
	}
	return a
}

// ContainsAllOf verifies that the actual string contains all of the given substrings.
//
// Examples:
//
//	assert.ThatString(t, "Gandalf the grey").ContainsAllOf("Gandalf", "grey") // OK
//	assert.ThatString(t, "Gandalf the grey").Contains("Gandalf", "white") // FAIL
func (a *StringAssert) ContainsAllOf(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	foundMissing := false
	for _, value := range values {
		if !strings.Contains(a.actual, value) {
			foundMissing = true
			break
		}
	}
	if foundMissing {
		a.failWithMessage("expected string to contain all of %#v, but got %q", values, a.actual)
	}
	return a
}

// ContainsAnyOf verifies that the actual string contains any of the given substrings.
//
// Example:
//
//	assert.ThatString(t, "Gandalf the Gray").ContainsAnyOf("Gandalf", "Saruman") // OK
//	assert.ThatString(t, "Bilbo Baggins").ContainsAnyOf("Frodo") // FAIL
func (a *StringAssert) ContainsAnyOf(values ...string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	foundAny := false
	for _, value := range values {
		if strings.Contains(a.actual, value) {
			foundAny = true
			break
		}
	}
	if !foundAny {
		a.failWithMessage("expected string to contain any of %#v, but got %q", values, a.actual)
	}
	return a
}

// DoesNotContain verifies that the actual string does not contain the given substring.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotContain("Pippin") // OK
//	assert.ThatString(t, "Frodo").DoesNotContain("do") // FAIL
func (a *StringAssert) DoesNotContain(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.Contains(a.actual, substr) {
		a.failWithMessage("expected string to not contain %q, but got %q", substr, a.actual)
	}
	return a
}

// DoesNotContainIgnoringCase verifies that the actual string does not contain the given substring, ignoring case considerations.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotContainIgnoringCase("sam") // OK
//	assert.ThatString(t, "Frodo").DoesNotContainIgnoringCase("fro") // FAIL
func (a *StringAssert) DoesNotContainIgnoringCase(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.Contains(strings.ToLower(a.actual), strings.ToLower(substr)) {
		a.failWithMessage("expected string to not contain %q ignoring case, but got %q", substr, a.actual)
	}
	return a
}

// DoesNotContainIgnoringWhitespace verifies that the actual string does not contain the given substring, ignoring whitespace.
//
// Example:
//
//	assert.ThatString(t, "Gandalf the grey").DoesNotContainIgnoringWhitespace("TheGrey") // OK
//	assert.ThatString(t, "Gandalf the grey").DoesNotContainIgnoringWhitespace("thegrey") // FAIL
func (a *StringAssert) DoesNotContainIgnoringWhitespace(substr string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.Contains(removeWhitespace(a.actual), removeWhitespace(substr)) {
		a.failWithMessage("expected string to not contain %q ignoring whitespace, but got %q", substr, a.actual)
	}
	return a
}

// StartsWith verifies that the actual string starts with the given prefix.
//
// Example:
//
//	assert.ThatString(t, "Frodo").StartsWith("Fro") // OK
//	assert.ThatString(t, "Frodo").StartsWith("fro") // FAIL
func (a *StringAssert) StartsWith(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.HasPrefix(a.actual, prefix) {
		a.failWithMessage("expected string to start with %q, but got %q", prefix, a.actual)
	}
	return a
}

// DoesNotStartWith verifies that the actual string does not start with the given prefix.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotStartWith("fro") // OK
//	assert.ThatString(t, "Frodo").DoesNotStartWith("") // FAIL
func (a *StringAssert) DoesNotStartWith(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.HasPrefix(a.actual, prefix) {
		a.failWithMessage("expected string to not start with %q, but got %q", prefix, a.actual)
	}
	return a
}

// StartsWithIgnoringCase verifies that the actual string starts with the given prefix, ignoring case considerations.
//
// Example:
//
//	assert.ThatString(t, "Frodo").StartsWith("fro") // OK
//	assert.ThatString(t, "Frodo").StartsWith("gan") // FAIL
func (a *StringAssert) StartsWithIgnoringCase(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.HasPrefix(strings.ToLower(a.actual), strings.ToLower(prefix)) {
		a.failWithMessage("expected string to start with %q ignoring case, but got %q", prefix, a.actual)
	}
	return a
}

// DoesNotStartWithIgnoringCase verifies that the actual string does not start with the given prefix, ignoring case considerations.
//
// Example:
//
//	assert.ThatString(t, "Gandalf").DoesNotStartWithIgnoringCase("Fro") // OK
//	assert.ThatString(t, "Gandalf").DoesNotStartWithIgnoringCase("gan") // FAIL
func (a *StringAssert) DoesNotStartWithIgnoringCase(prefix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.HasPrefix(strings.ToLower(a.actual), strings.ToLower(prefix)) {
		a.failWithMessage("expected string to not start with %q ignoring case, but got %q", prefix, a.actual)
	}
	return a
}

// EndsWith verifies that the actual string ends with the given suffix.
//
// Example:
//
//	assert.ThatString(t, "Frodo").EndsWith("do") // OK
//	assert.ThatString(t, "Frodo").EndsWith("Fro") // FAIL
func (a *StringAssert) EndsWith(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.HasSuffix(a.actual, suffix) {
		a.failWithMessage("expected string to end with %q, but got %q", suffix, a.actual)
	}
	return a
}

// DoesNotEndWith verifies that the actual string does not end with the given suffix.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotEndWith("Fro") // OK
//	assert.ThatString(t, "Frodo").DoesNotEndWith("do") // FAIL
func (a *StringAssert) DoesNotEndWith(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.HasSuffix(a.actual, suffix) {
		a.failWithMessage("expected string not to end with %q, but got %q", suffix, a.actual)
	}
	return a
}

// EndsWithIgnoringCase verifies that the actual string ends with the given suffix, ignoring case considerations.
//
// Example:
//
//	assert.ThatString(t, "Frodo").EndsWithIgnoringCase("Do") // OK
//	assert.ThatString(t, "Frodo").EndsWithIgnoringCase("Fro") // FAIL
func (a *StringAssert) EndsWithIgnoringCase(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.HasSuffix(strings.ToLower(a.actual), strings.ToLower(suffix)) {
		a.failWithMessage("expected string to end with %q ignoring case, but got %q", suffix, a.actual)
	}
	return a
}

// DoesNotEndWithIgnoringCase verifies that the actual string does not end with the given suffix, ignoring case considerations.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotEndWithIgnoringCase("Fro") // OK
//	assert.ThatString(t, "Frodo").DoesNotEndWithIgnoringCase("Do") // FAIL
func (a *StringAssert) DoesNotEndWithIgnoringCase(suffix string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if strings.HasSuffix(strings.ToLower(a.actual), strings.ToLower(suffix)) {
		a.failWithMessage("expected string not to end with %q ignoring case, but got %q", suffix, a.actual)
	}
	return a
}

// Matches verifies that the actual string matches the given regular expression.
//
// Example:
//
//	assert.ThatString(t, "Frodo").MatchesString(`..o.o`) // OK
//	assert.ThatString(t, "Frodo").MatchesString(`.*f`) // FAIL
func (a *StringAssert) Matches(pattern string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(a.actual) {
		a.failWithMessage("expected string to match %q, but got %q", pattern, a.actual)
	}
	return a
}

// DoesNotMatch verifies that the actual string does not match the given regular expression.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotMatch(`.*d$`) // OK
//	assert.ThatString(t, "Frodo").DoesNotMatch(`F.*`) // FAIL
func (a *StringAssert) DoesNotMatch(pattern string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	regex := regexp.MustCompile(pattern)
	if regex.MatchString(a.actual) {
		a.failWithMessage("expected string not to match %q, but got %q", pattern, a.actual)
	}
	return a
}

// MatchesRegexp verifies that the actual string matches the given compiled regular expression.
//
// Example:
//
//	assert.ThatString(t, "Frodo").MatchesString(regexp.MustCompile(`..o.o`)) // OK
//	assert.ThatString(t, "Frodo").MatchesString(regexp.MustCompile(`.*f`)) // FAIL
func (a *StringAssert) MatchesRegexp(regex *regexp.Regexp) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !regex.MatchString(a.actual) {
		a.failWithMessage("expected string to match \"%v\", but got %q", regex, a.actual)
	}
	return a
}

// DoesNotMatchRegexp verifies that the actual string does not match the given compiled regular expression.
//
// Example:
//
//	assert.ThatString(t, "Frodo").DoesNotMatchRegexp(regexp.MustCompile(`.*f`)) // OK
//	assert.ThatString(t, "Frodo").DoesNotMatchRegexp(regexp.MustCompile(`..o.o`)) // FAIL
func (a *StringAssert) DoesNotMatchRegexp(regex *regexp.Regexp) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if regex.MatchString(a.actual) {
		a.failWithMessage("expected string to match \"%v\", but got %q", regex, a.actual)
	}
	return a
}

// IsSubstringOf verifies that the actual string is a substring of the given string.
//
// Example:
//
//	assert.ThatString(t, "Lego").IsSubstringOf("Legolas") // OK
//	assert.ThatString(t, "Frodo").IsSubstringOf("Fro") // FAIL
func (a *StringAssert) IsSubstringOf(str string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !strings.Contains(str, a.actual) {
		a.failWithMessage("expected string to be a substring of %q, but got %q", str, a.actual)
	}
	return a
}

// IsIn verifies that the actual string is present in the given slice.
//
// Example:
//
//	assert.ThatString(t, "nenya").IsIn([]string{"vilya", "nenya", "narya"}) // OK
//	assert.ThatString(t, "one").IsIn([]string{"vilya", "nenya", "narya"}) // FAIL
func (a *StringAssert) IsIn(slice []string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if !slices.Contains(slice, a.actual) {
		a.failWithMessage("expected string to be present in %#v, but got %q", slice, a.actual)
	}
	return a
}

// IsNotIn verifies that the actual string is not present in the given slice.
//
// Example:
//
//	assert.ThatString(t, "one").IsNotIn([]string{"vilya", "nenya", "narya"}) // OK
//	assert.ThatString(t, "nenya").IsNotIn([]string{"vilya", "nenya", "narya"}) // FAIL
func (a *StringAssert) IsNotIn(slice []string) *StringAssert {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	if slices.Contains(slice, a.actual) {
		a.failWithMessage("expected string not to be present in %#v, but got %q", slice, a.actual)
	}
	return a
}
