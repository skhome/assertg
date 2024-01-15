package assert

import "fmt"

// DescriptionFormatter formats a description to be included in assertion errors.
type DescriptionFormatter func(description Description) string

// DefaultDescriptionFormatter formats a description by surrounding it in square brackets.
func DefaultDescriptionFormatter(description Description) string {
	if description != "" {
		return fmt.Sprintf("[%s] ", description)
	}
	return string(description)
}

//nolint:gochecknoglobals
var messageFormatter MessageFormatter = NewCompactMessageFormatter()

// MessageFormatter formats the message to be included in an assertion error.
type MessageFormatter interface {
	Format(description Description, representation Representation, format string, args ...any) string
}

// CompactMessageFormatter produces a compact assertion error message.
type CompactMessageFormatter struct {
	descriptionFormatter DescriptionFormatter
}

// NewCompactMessageFormatter creates and returns a new CompactMessageFormatter
func NewCompactMessageFormatter() CompactMessageFormatter {
	return CompactMessageFormatter{DefaultDescriptionFormatter}
}

func (f CompactMessageFormatter) Format(description Description, representation Representation, format string, args ...any) string {
	des := f.descriptionFormatter(description)
	msg := fmt.Sprintf(format, f.formatArgs(representation, args...)...)
	return fmt.Sprintf("%s%s", des, msg)
}

func (f CompactMessageFormatter) formatArgs(representation Representation, args ...any) []any {
	var formatted []any
	for _, arg := range args {
		formatted = append(formatted, f.asText(representation, arg))
	}
	return formatted
}

func (f CompactMessageFormatter) asText(representation Representation, value any) string {
	return representation(value)
}
