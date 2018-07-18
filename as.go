// Package as provides a easy way to convert numeric types with overflow check.
package as

import (
	"fmt"
)

// InvalidTypeError is returned when user supplied invalid type to convert.
type InvalidTypeError struct {
	ToType string
	Value  interface{}
}

func (e InvalidTypeError) Error() string {
	return fmt.Sprintf("unsupported type %T (converting to %s)", e.Value, e.ToType)
}

// OverflowError is returned when given value overflows maximum number a type can hold.
type OverflowError struct {
	ToType string
	Value  interface{}
}

func (e OverflowError) Error() string {
	return fmt.Sprintf("%d (%T) overflows %s", e.Value, e.Value, e.ToType)
}
