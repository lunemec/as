// +build 386 arm

package as

import (
	"math"
)

// Uint converts any numeric type to uint, and returns
// error if there was overflow.
func Uint(v interface{}) (uint, error) {
	var err error
	v = indirect(v)

	switch n := v.(type) {
	case int8:
		if n < 0 {
			err = OverflowError{ToType: "uint", Value: v}
		}
		return uint(n), err
	case int16:
		if n < 0 {
			err = OverflowError{ToType: "uint", Value: v}
		}
		return uint(n), err
	case int32:
		if n < 0 {
			err = OverflowError{ToType: "uint", Value: v}
		}
		return uint(n), err
	case int64:
		if n < 0 || n > math.MaxUint32 {
			err = OverflowError{ToType: "uint", Value: v}
		}
		return uint(n), err
	case int:
		if n < 0 {
			err = OverflowError{ToType: "uint", Value: v}
		}
		return uint(n), err
	case uint8:
		return uint(n), err
	case uint16:
		return uint(n), err
	case uint32:
		return uint(n), err
	case uint64:
		if n > math.MaxUint32 {
			err = OverflowError{ToType: "uint", Value: v}
		}
		return uint(n), err
	case uint:
		return uint(n), err
	}

	return 0, InvalidTypeError{ToType: "uint", Value: v}
}
