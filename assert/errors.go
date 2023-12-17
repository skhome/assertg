package assert

import "strings"

// tHelper is a helper interface to signal that this function is a test helper.
type tHelper interface {
	Helper()
}

// tName is a helper interface to retrieve the test name.
type tName interface {
	Name() string
}

// labeledContent contains labeled details of an assertion failure.
type labeledContent struct {
	label   string
	content string
}

// labeledOutput returns a string containing the given content as a table.
func labeledOutput(content ...labeledContent) string {
	longestLabel := 0
	for _, v := range content {
		if len(v.label) > longestLabel {
			longestLabel = len(v.label)
		}
	}
	var output string
	for _, v := range content {
		output += strings.Repeat(" ", longestLabel-len(v.label)) + v.label + ": " + v.content + "\n"
	}
	return output
}

// Fail reports a test failure.
func Fail(t TestingT, failureMessage string, description string) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	content := []labeledContent{
		{"Error", failureMessage},
	}
	if n, ok := t.(tName); ok {
		content = append(content, labeledContent{"Test", n.Name()})
	}
	if len(description) > 0 {
		content = append(content, labeledContent{"Description", description})
	}
	t.Errorf("\n%s", labeledOutput(content...))
}
