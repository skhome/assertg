package assert

import (
	"fmt"
)

// Supplier lazyly supplies values for an assertion info.
type Supplier[T any] func() T

// Description is a description for an assertion.
type Description string

// AssertionInfo provides information about an assertion.
type AssertionInfo interface {
	// OverridingFailureMessage returns the message that, if specified, will
	// replace the default message of an assertion failure.
	OverridingFailureMessage() string

	// Description returns the description of an assertion.
	Description() Description

	// Representation returns the Representation of the actual and expected values.
	Representation() Representation
}

// WritableAssertionInfo implements a modifiable AssertionInfo.
type WritableAssertionInfo struct {
	overridingFailureMessage         string
	overridingFailureMessageSupplier Supplier[string]
	description                      Description
	representation                   Representation
}

func NewWritableAssertionInfo() *WritableAssertionInfo {
	return &WritableAssertionInfo{
		representation: DefaultRepresentation,
	}
}

// OverridingFailureMessage returns the message that, if specified, will
// replace the default message of an assertion failure.
func (i *WritableAssertionInfo) OverridingFailureMessage() string {
	if i.overridingFailureMessageSupplier != nil {
		return i.overridingFailureMessageSupplier()
	}
	return i.overridingFailureMessage
}

// WithOverridingFailureMessage sets the failure message that replaces the default failure message.
func (i *WritableAssertionInfo) WithOverridingFailureMessage(message string, args ...any) {
	if len(args) > 0 {
		i.overridingFailureMessage = fmt.Sprintf(message, args...)
	} else {
		i.overridingFailureMessage = message
	}
}

// WithOverridingFailureMessageSupplier sets a lazy supplier for the failure message that replaces the default failure message.
func (i *WritableAssertionInfo) WithOverridingFailureMessageSupplier(supplier Supplier[string]) {
	i.overridingFailureMessageSupplier = supplier
}

// Description returns the description of an assertion.
func (i *WritableAssertionInfo) Description() Description {
	return i.description
}

// HasDescription returns if the assertion has a description.
func (i *WritableAssertionInfo) HasDescription() bool {
	return i.description != ""
}

// WithDescription sets a description for an assertion.
func (i *WritableAssertionInfo) WithDescription(description Description) {
	i.description = description
}

// Representation returns the Representation of actual and expected values.
func (i *WritableAssertionInfo) Representation() Representation {
	return i.representation
}

// UsingHexadecimalRepresentation uses a hexadecimal representation for actual and expected values.
func (i *WritableAssertionInfo) UsingHexadecimalRepresentation() {
	i.representation = HexadecimalRepresentation
}

// UsingBinaryRepresentation uses binary representation for actual and expected values.
func (i *WritableAssertionInfo) UsingBinaryRepresentation() {
	i.representation = BinaryRepresentation
}

// UsingRepresentation uses the given representation for actual and expected values.
func (i *WritableAssertionInfo) UsingRepresentation(representation Representation) {
	i.representation = representation
}
