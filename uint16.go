package as

import (
	"math"
)

// Uint16 converts any numeric type to uint16, and returns
// error if there was overflow.
func Uint16(v interface{}) (uint16, error) {
	var err error
	v = indirect(v)

	switch n := v.(type) {
	case int8:
		if n < 0 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case int16:
		if n < 0 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case int32:
		if n < 0 || n > math.MaxUint16 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case int64:
		if n < 0 || n > math.MaxUint16 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case int:
		if n < 0 || n > math.MaxUint16 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case uint8:
		return uint16(n), err
	case uint16:
		return n, err
	case uint32:
		if n > math.MaxUint16 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case uint64:
		if n > math.MaxUint16 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	case uint:
		if n > math.MaxUint16 {
			err = OverflowError{ToType: "uint16", Value: v}
		}
		return uint16(n), err
	}

	return 0, InvalidTypeError{ToType: "uint16", Value: v}
}
