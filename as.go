// Package as provides a easy way to convert numeric types with overflow check.
package as

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
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

// Number constraint is used for type cast into any number.
// Only Integers for now until float support is added.
type Number interface {
	constraints.Integer
}

// T is generic function to allow for easier checked
// type cast from any type to any Number type.
func T[To Number](v any) (To, error) {
	var (
		err error
		out To
	)
	switch any(out).(type) {
	case int:
		var n int
		n, err = Int(any(v))
		out = To(n)
	case int8:
		var n int8
		n, err = Int8(any(v))
		out = To(n)
	case int16:
		var n int16
		n, err = Int16(any(v))
		out = To(n)
	case int32:
		var n int32
		n, err = Int32(any(v))
		out = To(n)
	case int64:
		var n int64
		n, err = Int64(any(v))
		out = To(n)
	case uint:
		var n uint
		n, err = Uint(any(v))
		out = To(n)
	case uint8:
		var n uint8
		n, err = Uint8(any(v))
		out = To(n)
	case uint16:
		var n uint16
		n, err = Uint16(any(v))
		out = To(n)
	case uint32:
		var n uint32
		n, err = Uint32(any(v))
		out = To(n)
	case uint64:
		var n uint64
		n, err = Uint64(any(v))
		out = To(n)
	default:
		return out, InvalidTypeError{
			ToType: fmt.Sprintf("%T", out),
			Value:  v,
		}
	}
	return To(out), err
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
