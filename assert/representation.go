package assert

import "fmt"

// Representation provides a textual representation of a value in a certain format.
type Representation func(value any) string

// Returns the default representation of a given value.
func DefaultRepresentation(value any) string {
	return fmt.Sprintf("<%v>", value)
}

// HexadecimalRepresentation returns the hexadecimal representation of a given integer value.
func HexadecimalRepresentation(value any) string {
	return fmt.Sprintf("<%X>", value)
}

// BinaryRepresentation returns the binary representation of a given integer value.
func BinaryRepresentation(value any) string {
	return fmt.Sprintf("<%b>", value)
}
