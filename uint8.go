package as

import (
	"math"
)

// Uint8 converts any numeric type to uint8, and returns
// error if there was overflow.
func Uint8(v interface{}) (uint8, error) {
	var err error
	v = indirect(v)

	switch n := v.(type) {
	case int8:
		if n < 0 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case int16:
		if n < 0 || n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case int32:
		if n < 0 || n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case int64:
		if n < 0 || n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case int:
		if n < 0 || n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case uint8:
		return n, err
	case uint16:
		if n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case uint32:
		if n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case uint64:
		if n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	case uint:
		if n > math.MaxUint8 {
			err = OverflowError{ToType: "uint8", Value: v}
		}
		return uint8(n), err
	}

	return 0, InvalidTypeError{ToType: "uint8", Value: v}
}
