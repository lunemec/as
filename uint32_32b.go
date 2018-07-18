// +build 386 arm

package as

import (
	"math"
)

// Uint32 converts any numeric type to uint32, and returns
// error if there was overflow.
func Uint32(v interface{}) (uint32, error) {
	var err error

	switch n := v.(type) {
	case int8:
		if n < 0 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	case int16:
		if n < 0 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	case int32:
		if n < 0 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	case int64:
		if n < 0 || n > math.MaxUint32 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	case int:
		if n < 0 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	case uint8:
		return uint32(n), err
	case uint16:
		return uint32(n), err
	case uint32:
		return uint32(n), err
	case uint64:
		if n > math.MaxUint32 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	case uint:
		if n > math.MaxUint32 {
			err = OverflowError{ToType: "uint32", Value: v}
		}
		return uint32(n), err
	}

	return 0, InvalidTypeError{ToType: "uint32", Value: v}
}
