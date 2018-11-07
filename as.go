// Package as provides a easy way to convert numeric types with overflow check.
package as

import (
	"fmt"
	"reflect"
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
