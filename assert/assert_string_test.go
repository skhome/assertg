package assert_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/skhome/assertg/assert"
)

type stringTest struct {
	value   string
	other   string
	others  []string
	pattern string
	re      *regexp.Regexp
	num     int
	ok      bool
}

func TestStringIsEmpty(t *testing.T) {
	tests := []stringTest{
		{value: "", ok: true},
		{value: " ", ok: false},
		{value: "Frodo", ok: false},
	}
	messageFormat := "expected string to be empty, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsEmpty()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringIsNotEmpty(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", ok: true},
		{value: " ", ok: true},
		{value: "", ok: false},
	}
	messageFormat := "expected string to not be empty, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsNotEmpty()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringIsBlank(t *testing.T) {
	tests := []stringTest{
		{value: "", ok: true},
		{value: " ", ok: true},
		{value: " \t", ok: true},
		{value: "a", ok: false},
	}
	messageFormat := "expected string to be blank, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsBlank()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringIsNotBlank(t *testing.T) {
	tests := []stringTest{
		{value: "a", ok: true},
		{value: "", ok: false},
		{value: " ", ok: false},
		{value: "\t", ok: false},
	}
	messageFormat := "expected string to not be blank, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsNotBlank()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringContainsWhitespace(t *testing.T) {
	tests := []stringTest{
		{value: " ", ok: true},
		{value: "Frodo Baggins", ok: true},
		{value: "", ok: false},
		{value: "Frodo", ok: false},
	}
	messageFormat := "expected string to contain whitespace characters, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsWhitespace()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringDoesNotContainWhitespace(t *testing.T) {
	tests := []stringTest{
		{value: "", ok: true},
		{value: "Frodo", ok: true},
		{value: " ", ok: false},
		{value: "Frodo Baggins", ok: false},
	}
	messageFormat := "expected string to not contain whitespace characters, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotContainWhitespace()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringHasLength(t *testing.T) {
	tests := []stringTest{
		{value: "Lord of the Rings", num: 17, ok: true},
		{value: "", num: 0, ok: true},
		{value: "foo", num: 2, ok: false},
	}
	messageFormat := "expected string to have length of <%d>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).HasLength(test.num)
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.value)
	})
}

func TestStringHasLengthLessThan(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", num: 6, ok: true},
		{value: "Frodo", num: 5, ok: false},
	}
	messageFormat := "expected string to have length less than <%d>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).HasLengthLessThan(test.num)
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.value)
	})
}

func TestStringHasLengthGreaterThan(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", num: 4, ok: true},
		{value: "Frodo", num: 5, ok: false},
	}
	messageFormat := "expected string to have length greater than <%d>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).HasLengthGreaterThan(test.num)
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.value)
	})
}

func TestStringHasSameLengthAs(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf", other: "Saruman", ok: true},
		{value: "", other: "", ok: true},
		{value: "Gandalf", other: "Frodo", ok: false},
	}
	messageFormat := "expected string to have the same length as <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).HasSameLengthAs(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringHasLineCount(t *testing.T) {
	tests := []stringTest{
		{value: "first", num: 1, ok: true},
		{value: "", num: 1, ok: true},
		{value: "first\nsecond", num: 2, ok: true},
		{value: "first\nsecond\n", num: 3, ok: true},
		{value: "first", num: 2, ok: false},
	}
	messageFormat := "expected string to have <%d> lines, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).HasLineCount(test.num)
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.value)
	})
}

func TestStringIsEqualTo(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "Frodo", ok: true},
		{value: "Frodo", other: "Sam", ok: false},
	}
	messageFormat := "expected string to equal <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringIsNotEqualTo(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "frodo", ok: true},
		{value: "Frodo", other: "Frodo", ok: false},
	}
	messageFormat := "expected string not to equal <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsNotEqualTo(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringIsEqualToIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "frodo", ok: true},
		{value: "Frodo", other: "gandalf", ok: false},
	}
	messageFormat := "expected string to equal <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsEqualToIgnoringCase(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringIsNotEqualToIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "gandalf", ok: true},
		{value: "Frodo", other: "frodo", ok: false},
	}
	messageFormat := "expected string to not equal <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsNotEqualToIgnoringCase(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringContainsDigit(t *testing.T) {
	tests := []stringTest{
		{value: "Bug8ear", ok: true},
		{value: "V", ok: false},
	}
	messageFormat := "expected string to contain digit, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsDigit()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringContainsOnlyDigits(t *testing.T) {
	tests := []stringTest{
		{value: "10", ok: true},
		{value: "10a", ok: false},
	}
	messageFormat := "expected string to only contain digits, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsOnlyDigits()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringContainsOnlyOnce(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "do", ok: true},
		{value: "Frodo", other: "o", ok: false},
		{value: "Frodo", other: "y", ok: false},
	}
	messageFormat := "expected string to contain <%s> only once, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsOnlyOnce(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringContains(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", others: []string{"alf", "grey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"white"}, ok: false},
	}
	messageFormat := "expected string to contain <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).Contains(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringContainsAnyOf(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", others: []string{"grey", "black"}, ok: true},
		{value: "Gandalf the grey", others: []string{"white", "black"}, ok: false},
	}
	messageFormat := "expected string to contain any of <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsAnyOf(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringContainsIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", others: []string{"gandalf", "Grey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"white"}, ok: false},
	}
	messageFormat := "expected string to contain <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsIgnoringCase(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringContainsIgnoringWhitespace(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", others: []string{"alf"}, ok: true},
		{value: "Gandalf the grey", others: []string{"alf", "grey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"thegrey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"thegr  ey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"t h e\t grey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"alF"}, ok: false},
	}
	messageFormat := "expected string to contain <%s> ignoring whitespace, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).ContainsIgnoringWhitespace(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringDoesNotContain(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", others: []string{"pippin"}, ok: true},
		{value: "Frodo", others: []string{"fro", "sam"}, ok: true},
		{value: "Frodo", others: []string{"Fro", "Gimli"}, ok: false},
	}
	messageFormat := "expected string to not contain <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotContain(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringDoesNotContainIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", others: []string{"pippin"}, ok: true},
		{value: "Frodo", others: []string{"Merry", "sam"}, ok: true},
		{value: "Frodo", others: []string{"Fro", "Gimli"}, ok: false},
		{value: "Frodo", others: []string{"fro"}, ok: false},
	}
	messageFormat := "expected string to not contain <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotContainIgnoringCase(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringDoesNotContainIgnoringWhitespace(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", others: []string{"TheGrey"}, ok: true},
		{value: "Gandalf the grey", others: []string{"thegrey"}, ok: false},
	}
	messageFormat := "expected string to not contain <%s> ignoring whitespace, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotContainIgnoringWhitespace(test.others...)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringStartsWith(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "Fro", ok: true},
		{value: "Frodo", other: "fro", ok: false},
	}
	messageFormat := "expected string to start with <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).StartsWith(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringDoesNotStartWith(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "fro", ok: true},
		{value: "Frodo", other: "Fro", ok: false},
	}
	messageFormat := "expected string not to start with <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotStartWith(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringStartsWithIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", other: "Gandalf", ok: true},
		{value: "Gandalf the grey", other: "gandalf", ok: true},
		{value: "Gandald the grey", other: "grey", ok: false},
	}
	messageFormat := "expected string to start with <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).StartsWithIgnoringCase(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringDoesNotStartWithIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Gandalf the grey", other: "fro", ok: true},
		{value: "Gandalf the grey", other: "grey", ok: true},
		{value: "Gandalf the grey", other: "Gandalf", ok: false},
		{value: "Gandalf the grey", other: "gandalf", ok: false},
	}
	messageFormat := "expected string not to start with <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotStartWithIgnoringCase(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringEndsWith(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "do", ok: true},
		{value: "Frodo", other: "Fro", ok: false},
	}
	messageFormat := "expected string to end with <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).EndsWith(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringDoesNotEndWith(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "Fro", ok: true},
		{value: "Frodo", other: "do", ok: false},
		{value: "Frodo", other: "", ok: false},
	}
	messageFormat := "expected string not to end with <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotEndWith(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringEndsWithIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "do", ok: true},
		{value: "Frodo", other: "Do", ok: true},
		{value: "Frodo", other: "fro", ok: false},
	}
	messageFormat := "expected string to end with <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).EndsWithIgnoringCase(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringDoesNotEndWithIgnoringCase(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", other: "Fro", ok: true},
		{value: "Frodo", other: "do", ok: false},
		{value: "Frodo", other: "DO", ok: false},
	}
	messageFormat := "expected string not to end with <%s> ignoring case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotEndWithIgnoringCase(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringMatchesPattern(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", pattern: `..o.o`, ok: true},
		{value: "Frodo", pattern: `.*d$`, ok: false},
	}
	messageFormat := "expected string to match <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).MatchesPattern(test.pattern)
		return test.ok, fmt.Sprintf(messageFormat, test.pattern, test.value)
	})
}

func TestStringDoesNotMatchPattern(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", pattern: `.*d$`, ok: true},
		{value: "Frodo", pattern: `..o.o`, ok: false},
	}
	messageFormat := "expected string not to match <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotMatchPattern(test.pattern)
		return test.ok, fmt.Sprintf(messageFormat, test.pattern, test.value)
	})
}

func TestStringMatchesRegexp(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", re: regexp.MustCompile(`..o.o`), ok: true},
		{value: "Frodo", re: regexp.MustCompile(`.*d$`), ok: false},
	}
	messageFormat := "expected string to match <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).MatchesRegexp(test.re)
		return test.ok, fmt.Sprintf(messageFormat, test.re, test.value)
	})
}

func TestStringDoesNotMatchRegexp(t *testing.T) {
	tests := []stringTest{
		{value: "Frodo", re: regexp.MustCompile(`.*d$`), ok: true},
		{value: "Frodo", re: regexp.MustCompile(`..o.o`), ok: false},
	}
	messageFormat := "expected string to match <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).DoesNotMatchRegexp(test.re)
		return test.ok, fmt.Sprintf(messageFormat, test.re, test.value)
	})
}

func TestStringIsEqualToIgnoringWhitespace(t *testing.T) {
	tests := []stringTest{
		{value: "Game of Thrones", other: "Game   of   Thrones", ok: true},
		{value: "Game of Thrones", other: "  Game of   Thrones  ", ok: true},
		{value: "Game of Thrones", other: "  Game of Thrones  ", ok: true},
		{value: "Game of Thrones", other: "\tGame of Thrones\n", ok: true},
		{value: "Game of Thrones", other: "Game OF Thrones", ok: false},
	}
	messageFormat := "expected string to equal ignoring whitespace <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsEqualToIgnoringWhitespace(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringIsNotEqualToIgnoringWhitespace(t *testing.T) {
	tests := []stringTest{
		{value: "Game of Thrones", other: "Game OF Thrones", ok: true},
		{value: "Game of Thrones", other: "Game   of   Thrones", ok: false},
		{value: "Game of Thrones", other: "  Game of   Thrones  ", ok: false},
		{value: "Game of Thrones", other: "  Game of Thrones  ", ok: false},
		{value: "Game of Thrones", other: "\tGame of Thrones\n", ok: false},
	}
	messageFormat := "expected string not to equal ignoring whitespace <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsNotEqualToIgnoringWhitespace(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringIsSubstringOf(t *testing.T) {
	tests := []stringTest{
		{value: "Lego", other: "Legolas", ok: true},
		{value: "Lego", other: "Lego", ok: true},
		{value: "Frodo", other: "Frod", ok: false},
	}
	messageFormat := "expected string to be a substring of <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsSubstringOf(test.other)
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.value)
	})
}

func TestStringIsIn(t *testing.T) {
	hobbits := []string{"Frodo", "Sam", "Merry", "Pippin", "Bilbo"}
	tests := []stringTest{
		{value: "Frodo", others: hobbits, ok: true},
		{value: "Legolas", others: hobbits, ok: false},
		{value: "Legolas", others: nil, ok: false},
	}
	messageFormat := "expected string to be present in <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsIn(test.others)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringIsNotIn(t *testing.T) {
	hobbits := []string{"Frodo", "Sam", "Merry", "Pippin", "Bilbo"}
	tests := []stringTest{
		{value: "Legolas", others: hobbits, ok: true},
		{value: "Frodo", others: nil, ok: true},
		{value: "Frodo", others: hobbits, ok: false},
	}
	messageFormat := "expected string not to be present in <%s>, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsNotIn(test.others)
		return test.ok, fmt.Sprintf(messageFormat, test.others, test.value)
	})
}

func TestStringLowerCase(t *testing.T) {
	tests := []stringTest{
		{value: "legolas", ok: true},
		{value: "", ok: true},
		{value: ".", ok: true},
		{value: "42", ok: true},
		{value: "Legolas", ok: false},
	}
	messageFormat := "expected string to be all lower case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsLowerCase()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}

func TestStringUpperCase(t *testing.T) {
	tests := []stringTest{
		{value: "LEGOLAS", ok: true},
		{value: "", ok: true},
		{value: ".", ok: true},
		{value: "42", ok: true},
		{value: "Legolas", ok: false},
	}
	messageFormat := "expected string to be all upper case, but got <%s>"
	runTests(t, tests)(func(fixture *fixtureT, test stringTest) (bool, string) {
		assert.ThatString(fixture, test.value).IsUpperCase()
		return test.ok, fmt.Sprintf(messageFormat, test.value)
	})
}
