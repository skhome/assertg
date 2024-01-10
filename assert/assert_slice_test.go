package assert_test

import (
	"fmt"
	"testing"

	"github.com/skhome/assertg/assert"
)

type sliceTest struct {
	slice   []string
	num     int
	other   []string
	ok      bool
	message string
}

func prepareSliceAssertWithMessage(fixture *fixtureT, test sliceTest) *assert.SliceAssert[string] {
	sliceAssert := assert.ThatSlice(fixture, test.slice)
	if test.message != "" {
		sliceAssert.WithFailMessage(test.message)
	}
	return sliceAssert
}

func TestSliceIsNil(t *testing.T) {
	tests := []sliceTest{
		{slice: []string(nil), ok: true},
		{slice: []string{}, ok: false},
		{slice: []string{"Frodo"}, ok: false},
		{slice: []string{}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to be nil, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.IsNil()
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceIsNotNil(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"Frodo"}, ok: true},
		{slice: []string{}, ok: true},
		{slice: []string(nil), ok: false},
		{slice: []string(nil), ok: false, message: "my message"},
	}
	messageFormat := "expected slice to not be nil, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.IsNotNil()
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceIsEmpty(t *testing.T) {
	tests := []sliceTest{
		{slice: []string(nil), ok: true},
		{slice: []string{}, ok: true},
		{slice: []string{"Hobbit"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to be empty, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.IsEmpty()
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceIsNotEmpty(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"Nenya", "Narya", "Vilya"}, ok: true},
		{slice: []string{}, ok: false},
		{slice: []string(nil), ok: false},
		{slice: []string{}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to not be empty, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.IsNotEmpty()
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceHasSize(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"Nenya", "Narya", "Vilya"}, num: 3, ok: true},
		{slice: []string{}, num: 0, ok: true},
		{slice: []string(nil), num: 0, ok: true},
		{slice: []string{}, num: 1, ok: false},
		{slice: []string(nil), num: 1, ok: false},
	}
	messageFormat := "expected slice to have a size of %d, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasSize(test.num)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.slice)
	})
}

func TestSliceHasSizeGreaterThan(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"Nenya", "Narya", "Vilya"}, num: 2, ok: true},
		{slice: []string{}, num: 0, ok: false},
		{slice: []string(nil), num: 0, ok: false},
		{slice: []string{}, num: 0, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to have a size greater than %d, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasSizeGreaterThan(test.num)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.slice)
	})
}

func TestSliceHasSizeLessThan(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"Nenya", "Narya", "Vilya"}, num: 4, ok: true},
		{slice: []string{"Frodo"}, num: 1, ok: false},
		{slice: []string{}, num: 0, ok: false},
		{slice: []string(nil), num: 0, ok: false},
		{slice: []string{"Frodo"}, num: 1, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to have a size less than %d, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasSizeLessThan(test.num)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.slice)
	})
}

func TestSliceHasSameSizeAs(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"Frodo", "Sam"}, other: []string{"Merry", "Pippin"}, ok: true},
		{slice: []string{}, other: []string{}, ok: true},
		{slice: []string(nil), other: []string{}, ok: true},
		{slice: []string{"Frodo", "Sam"}, other: []string{"Gandalf"}, ok: false},
		{slice: []string(nil), other: []string{"Frodo"}, ok: false},
		{slice: []string(nil), other: []string{"Frodo"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to have the same size as %#v (%d), but got %#v (%d)"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasSameSizeAs(test.other)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, len(test.other), test.slice, len(test.slice))
	})
}

func TestSliceContains(t *testing.T) {
	abc := []string{"a", "b", "c"}
	tests := []sliceTest{
		{slice: abc, other: []string{"a", "b"}, ok: true},
		{slice: abc, other: []string{"c", "a"}, ok: true},
		{slice: abc, other: []string{"a", "c", "b"}, ok: true},
		{slice: abc, other: []string{"a", "f"}, ok: false},
		{slice: abc, other: []string{"d"}, ok: false},
		{slice: nil, other: []string{"a"}, ok: false},
		{slice: abc, other: []string{"d"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to contain %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.Contains(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceContainsOnly(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "c"}, other: []string{"a", "b", "c"}, ok: true},
		{slice: []string{"a", "b", "c"}, other: []string{"b", "c", "a"}, ok: true},
		{slice: []string{"a", "a", "b"}, other: []string{"a", "b"}, ok: true},
		{slice: []string{}, other: []string{}, ok: true},
		{slice: nil, other: nil, ok: true},
		{slice: []string{"a", "b", "c"}, other: []string{"a", "b", "c", "d"}, ok: false},
		{slice: []string{"a", "b", "c"}, other: []string{"a", "b"}, ok: false},
		{slice: []string{"a", "b", "c"}, other: []string{"d", "e"}, ok: false},
		{slice: []string{"a", "b", "c"}, other: []string{"d", "e"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to contain only %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.ContainsOnly(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceContainsOnlyOnce(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "c"}, other: []string{"a", "b"}, ok: true},
		{slice: []string{"a", "b", "b"}, other: []string{"a"}, ok: true},
		{slice: []string{"a", "b"}, other: []string{"b", "a"}, ok: true},
		{slice: []string{}, other: []string{}, ok: true},
		{slice: nil, other: nil, ok: true},
		{slice: []string{"a", "b", "a"}, other: []string{"a"}, ok: false},
		{slice: []string{"a", "b", "c"}, other: []string{"d"}, ok: false},
		{slice: []string{"a", "b", "c", "c"}, other: []string{"c", "d"}, ok: false},
	}
	messageFormat := "expected slice to contain %#v only once, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.ContainsOnlyOnce(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceContainsExactly(t *testing.T) {
	elvenRings := []string{"vilya", "nenya", "narya"}
	tests := []sliceTest{
		{slice: elvenRings, other: []string{"vilya", "nenya", "narya"}, ok: true},
		{slice: []string{}, other: []string{}, ok: true},
		{slice: []string(nil), other: []string(nil), ok: true},
		{slice: elvenRings, other: []string{"vilya"}, ok: false},
		{slice: elvenRings, other: []string{"nenya", "vilya"}, ok: false},
		{slice: elvenRings, other: []string{}, ok: false},
		{slice: []string{}, other: []string{"nenya"}, ok: false},
		{slice: elvenRings, other: []string{"vilya"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to contain exactly %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.ContainsExactly(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceContainsExactlyInAnyOrder(t *testing.T) {
	elvenRings := []string{"vilya", "nenya", "narya", "vilya"}
	tests := []sliceTest{
		{slice: elvenRings, other: []string{"vilya", "vilya", "nenya", "narya"}, ok: true},
		{slice: elvenRings, other: []string{"vilya", "nenya", "narya"}, ok: false},
		{slice: elvenRings, other: []string{"vilya", "nenya", "narya"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to contain exactly %v in any order, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.ContainsExactlyInAnyOrder(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceContainsSequence(t *testing.T) {
	elvenRings := []string{"vilya", "nenya", "narya"}
	tests := []sliceTest{
		{slice: elvenRings, other: []string{"vilya", "nenya"}, ok: true},
		{slice: elvenRings, other: []string{"nenya", "narya"}, ok: true},
		{slice: elvenRings, other: []string{}, ok: true},
		{slice: elvenRings, other: []string{"one"}, ok: false},
		{slice: elvenRings, other: []string{"vilya", "narya"}, ok: false},
		{slice: elvenRings, other: []string{"nenya", "vilya"}, ok: false},
		{slice: elvenRings, other: []string{"nenya", "vilya"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to contain the sequence %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.ContainsSequence(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceDoesNotContainsSequence(t *testing.T) {
	elvenRings := []string{"vilya", "nenya", "narya"}
	tests := []sliceTest{
		{slice: elvenRings, other: []string{"vilya", "narya"}, ok: true},
		{slice: elvenRings, other: []string{"nenya", "vilya"}, ok: true},
		{slice: elvenRings, other: []string{"vilya", "nenya"}, ok: false},
		{slice: elvenRings, other: []string{"nenya", "narya"}, ok: false},
		{slice: elvenRings, other: []string{"nenya", "narya"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice not to contain the sequence %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.DoesNotContainSequence(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceDoesNotContain(t *testing.T) {
	abc := []string{"a", "b", "c"}
	tests := []sliceTest{
		{slice: abc, other: []string{"d"}, ok: true},
		{slice: abc, other: []string{"d", "e"}, ok: true},
		{slice: abc, other: []string{"a"}, ok: false},
		{slice: abc, other: []string{"a", "b"}, ok: false},
		{slice: abc, other: []string{"c", "d"}, ok: false},
		{slice: abc, other: []string{"a"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice not to contain %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.DoesNotContain(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceContainsAnyOf(t *testing.T) {
	abc := []string{"a", "b", "c"}
	tests := []sliceTest{
		{slice: abc, other: []string{"b"}, ok: true},
		{slice: abc, other: []string{"b", "c"}, ok: true},
		{slice: abc, other: []string{"a", "b", "c"}, ok: true},
		{slice: abc, other: []string{"a", "b", "c", "d"}, ok: true},
		{slice: abc, other: []string{"e", "f", "c"}, ok: true},
		{slice: abc, other: []string{"d"}, ok: false},
		{slice: abc, other: []string{"d", "e", "f", "g"}, ok: false},
		{slice: abc, other: []string{"d"}, ok: false, message: "my message"},
	}
	messageFormat := "expected slice to contain any of %#v, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.ContainsAnyOf(test.other...)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.other, test.slice)
	})
}

func TestSliceHasAll(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "c"}, ok: true},
		{slice: []string{"a", "b", "cc"}, ok: false},
		{slice: []string{"a", "b", "cc"}, ok: false, message: "my message"},
	}
	isSingleCharacter := func(value string) bool { return len(value) == 1 }
	messageFormat := "expected slice to have all entries match the predicate, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasAll(isSingleCharacter)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceHasNone(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "c"}, ok: true},
		{slice: []string{"a", "b", "cc"}, ok: false},
		{slice: []string{"a", "b", "cc"}, ok: false, message: "my message"},
	}
	isMultiCharacter := func(value string) bool { return len(value) > 1 }
	messageFormat := "expected slice to have no entry match the predicate, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasNone(isMultiCharacter)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceHasAny(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "bb", "cc"}, ok: true},
		{slice: []string{"a", "b", "cc"}, ok: true},
		{slice: []string{"aa", "bb", "cc"}, ok: false},
		{slice: []string{"aa", "bb", "cc"}, ok: false, message: "my message"},
	}
	isSingleCharacter := func(value string) bool { return len(value) == 1 }
	messageFormat := "expected slice to have any entry match the predicate, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasAny(isSingleCharacter)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.slice)
	})
}

func TestSliceHasAtLeast(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "cc"}, num: 2, ok: true},
		{slice: []string{"a", "b", "cc"}, num: 1, ok: true},
		{slice: []string{"a", "b", "cc"}, num: 3, ok: false},
		{slice: []string{"a", "b", "cc"}, num: 3, ok: false, message: "my message"},
	}
	isSingleCharacter := func(value string) bool { return len(value) == 1 }
	messageFormat := "expected slice to have at least %d entries match the predicate, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasAtLeast(test.num, isSingleCharacter)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.slice)
	})
}

func TestSliceHasAtMostMatch(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "cc"}, num: 2, ok: true},
		{slice: []string{"a", "b", "cc"}, num: 3, ok: true},
		{slice: []string{"a", "b", "cc"}, num: 1, ok: false},
		{slice: []string{"a", "b", "cc"}, num: 1, ok: false, message: "my message"},
	}
	isSingleCharacter := func(value string) bool { return len(value) == 1 }
	messageFormat := "expected slice to have at most %d entries match the predicate, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasAtMost(test.num, isSingleCharacter)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.slice)
	})
}

func TestSliceHasExactly(t *testing.T) {
	tests := []sliceTest{
		{slice: []string{"a", "b", "cc"}, num: 2, ok: true},
		{slice: []string{"a", "b", "cc"}, num: 1, ok: false},
		{slice: []string{"a", "b", "cc"}, num: 3, ok: false},
		{slice: []string{"a", "b", "cc"}, num: 1, ok: false, message: "my message"},
	}
	isSingleCharacter := func(value string) bool { return len(value) == 1 }
	messageFormat := "expected slice to have exactly %d entries match the predicate, but got %#v"
	runTests(t, tests)(func(fixture *fixtureT, test sliceTest) (bool, string) {
		sliceAssert := prepareSliceAssertWithMessage(fixture, test)
		sliceAssert.HasExactly(test.num, isSingleCharacter)
		if test.message != "" {
			return test.ok, test.message
		}
		return test.ok, fmt.Sprintf(messageFormat, test.num, test.slice)
	})
}

func TestSliceExtractingStrings(t *testing.T) {
	type TolkienCharacter struct {
		name    string
		age     int
		species string
	}
	fellowship := []TolkienCharacter{
		{name: "Frodo", age: 33, species: "Hobbit"},
		{name: "Sam", age: 38, species: "Hobbit"},
		{name: "Gandalf", age: 2020, species: "Maia"},
		{name: "Legolas", age: 1000, species: "Elf"},
	}
	characterName := func(character TolkienCharacter) string { return character.name }
	assert.ThatSlice(t, fellowship).
		ExtractingStrings(characterName).
		ContainsExactly("Frodo", "Sam", "Gandalf", "Legolas")
}

func TestSliceExtractingInts(t *testing.T) {
	type TolkienCharacter struct {
		name    string
		age     int
		species string
	}
	fellowship := []TolkienCharacter{
		{name: "Frodo", age: 33, species: "Hobbit"},
		{name: "Sam", age: 38, species: "Hobbit"},
		{name: "Gandalf", age: 2020, species: "Maia"},
		{name: "Legolas", age: 1000, species: "Elf"},
	}
	characterAge := func(character TolkienCharacter) int { return character.age }
	assert.ThatSlice(t, fellowship).
		ExtractingInts(characterAge).
		ContainsExactly(33, 38, 2020, 1000)
}

func TestSliceExtracting(t *testing.T) {
	type Species string
	type TolkienCharacter struct {
		name    string
		species Species
	}
	fellowship := []TolkienCharacter{
		{name: "Frodo", species: "Hobbit"},
		{name: "Legolas", species: "Elf"},
	}
	characterSpecies := func(character TolkienCharacter) any { return character.species }
	assert.ThatSlice(t, fellowship).
		Extracting(characterSpecies).
		Contains(Species("Hobbit"))
}
