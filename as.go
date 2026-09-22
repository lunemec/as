// Package as converts integer types and reports overflow.
package as

import (
	"fmt"
	"reflect"
)

// OverflowError is returned when a value cannot be represented by the target type.
type OverflowError struct {
	ToType string
	Value  any
}

func (e OverflowError) Error() string {
	return fmt.Sprintf("%d (%T) overflows %s", e.Value, e.Value, e.ToType)
}

// Number contains all supported integer types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// T converts v to To and reports whether the value overflows To.
func T[To, From Number](v From) (To, error) {
	out, ok := checkedCast[To](v)
	if !ok {
		return out, overflowError[To](v)
	}
	return out, nil
}

func checkedCast[To, From Number](v From) (To, bool) {
	out := To(v)
	ok := From(out) == v &&
		(isSigned[From]() == isSigned[To]() || (v >= 0 && out >= 0))
	return out, ok
}

func isSigned[T Number]() bool {
	return ^T(0) < 0
}

func overflowError[To, From Number](v From) error {
	return OverflowError{
		ToType: reflect.TypeFor[To]().String(),
		Value:  v,
	}
}

// Int converts v to int and reports overflow.
func Int[From Number](v From) (int, error) { return T[int](v) }

// Int8 converts v to int8 and reports overflow.
func Int8[From Number](v From) (int8, error) { return T[int8](v) }

// Int16 converts v to int16 and reports overflow.
func Int16[From Number](v From) (int16, error) { return T[int16](v) }

// Int32 converts v to int32 and reports overflow.
func Int32[From Number](v From) (int32, error) { return T[int32](v) }

// Int64 converts v to int64 and reports overflow.
func Int64[From Number](v From) (int64, error) { return T[int64](v) }

// Uint converts v to uint and reports overflow.
func Uint[From Number](v From) (uint, error) { return T[uint](v) }

// Uint8 converts v to uint8 and reports overflow.
func Uint8[From Number](v From) (uint8, error) { return T[uint8](v) }

// Uint16 converts v to uint16 and reports overflow.
func Uint16[From Number](v From) (uint16, error) { return T[uint16](v) }

// Uint32 converts v to uint32 and reports overflow.
func Uint32[From Number](v From) (uint32, error) { return T[uint32](v) }

// Uint64 converts v to uint64 and reports overflow.
func Uint64[From Number](v From) (uint64, error) { return T[uint64](v) }

// Uintptr converts v to uintptr and reports overflow.
func Uintptr[From Number](v From) (uintptr, error) { return T[uintptr](v) }
