package as

import (
	"math"
)

// Int8 converts any numeric type to int8, and returns
// error if there was overflow.
func Int8(v interface{}) (int8, error) {
	var err error
	v = indirect(v)

	switch n := v.(type) {
	case int8:
		return n, err
	case int16:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case int32:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case int64:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case int:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case uint8:
		if n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case uint16:
		if n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case uint32:
		if n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case uint64:
		if n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	case uint:
		if n > math.MaxInt8 {
			err = OverflowError{ToType: "int8", Value: v}
		}
		return int8(n), err
	}

	return 0, InvalidTypeError{ToType: "int8", Value: v}
}
