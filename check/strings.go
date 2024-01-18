package check

import (
	"regexp"
	"strings"
	"unicode"
)

// StringIsBlank returns if the string is empty or contains only whitespace characters.
func StringIsBlank(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// StringContainsWhitespace returns if the string contains any whitespace characters (as defined by unicode.IsSpace).
func StringContainsWhitespace(value string) bool {
	return strings.ContainsFunc(value, unicode.IsSpace)
}

// StringLineCount returns the number of lines in the given string.
func StringLineCount(value string) int {
	return strings.Count(value, "\n") + 1
}

// StringContainsDigit returns if the string contains any digit (as defined by unicode.IsDigit).
func StringContainsDigit(value string) bool {
	return strings.ContainsFunc(value, unicode.IsDigit)
}

// StringContainsOnlyDigits returns if the string contains only digits.
func StringContainsOnlyDigits(value string) bool {
	removeDigit := func(r rune) rune {
		if unicode.IsDigit(r) {
			return -1
		}
		return r
	}
	return len(strings.Map(removeDigit, value)) == 0
}

// StringContains returns if the string contains all given substrings.
func StringContains(s string, substrs []string) bool {
	for i := range substrs {
		if !strings.Contains(s, substrs[i]) {
			return false
		}
	}
	return true
}

// StringContainsAny returns if the string contains any of the given substrings.
func StringContainsAny(s string, substrs []string) bool {
	for i := range substrs {
		if strings.Contains(s, substrs[i]) {
			return true
		}
	}
	return false
}

// StringContainsIgnoringCase returns if the string contains all given substrings.
func StringContainsIgnoringCase(s string, substrs []string) bool {
	for i := range substrs {
		if !strings.Contains(strings.ToLower(s), strings.ToLower(substrs[i])) {
			return false
		}
	}
	return true
}

// StringContainsAnyIgnoringCase returns if the string contains any of the given substrings ignoring case.
func StringContainsAnyIgnoringCase(s string, substrs []string) bool {
	for i := range substrs {
		if strings.Contains(strings.ToLower(s), strings.ToLower(substrs[i])) {
			return true
		}
	}
	return false
}

// stringRemoveWhitespace returns a string with all whitespace stripped.
func stringRemoveWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func StringEqualsIgnoringWhitespace(s, o string) bool {
	return stringRemoveWhitespace(s) == stringRemoveWhitespace(o)
}

// StringContainsIgnoringWhitespace returns the string contains all given substrings ignoring whitespace.
func StringContainsIgnoringWhitespace(s string, substrs []string) bool {
	for i := range substrs {
		if !strings.Contains(stringRemoveWhitespace(s), stringRemoveWhitespace(substrs[i])) {
			return false
		}
	}
	return true
}

// StringContainsAnyIgnoringWhitespace returns if the string contains any of the given substrings ignoring whitespace.
func StringContainsAnyIgnoringWhitespace(s string, substrs []string) bool {
	for i := range substrs {
		if strings.Contains(stringRemoveWhitespace(s), stringRemoveWhitespace(substrs[i])) {
			return true
		}
	}
	return false
}

// StringStartsWith returns if the string starts with the given prefix.
func StringStartsWith(value string, prefix string) bool {
	return strings.HasPrefix(value, prefix)
}

// StringStartsWithIgnoringCase returns if the string starts with the given prefix ignoring case.
func StringStartsWithIgnoringCase(value string, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(value), strings.ToLower(prefix))
}

// StringEndsWith returns if the string ends with the given suffix.
func StringEndsWith(value string, suffix string) bool {
	return strings.HasSuffix(value, suffix)
}

// StringEndsWithIgnoringCase returns if the string ends with the given suffix ignoring case.
func StringEndsWithIgnoringCase(value string, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(value), strings.ToLower(suffix))
}

// StringMatchesRegexp returns if the string matches the given regular expression.
func StringMatchesRegexp(value string, re *regexp.Regexp) bool {
	return re.MatchString(value)
}

// StringIsEqual returns if both strings are equal.
func StringIsEqual(a, b string) bool {
	return a == b
}
