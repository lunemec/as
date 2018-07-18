package as

import (
	"math"
)

// Int32 converts any numeric type to int32, and returns
// error if there was overflow.
func Int32(v interface{}) (int32, error) {
	var err error

	switch n := v.(type) {
	case int8:
		return int32(n), err
	case int16:
		return int32(n), err
	case int32:
		return n, err
	case int64:
		if n < math.MinInt32 || n > math.MaxInt32 {
			err = OverflowError{ToType: "int32", Value: v}
		}
		return int32(n), err
	case int:
		if n < math.MinInt32 || n > math.MaxInt32 {
			err = OverflowError{ToType: "int32", Value: v}
		}
		return int32(n), err
	case uint8:
		return int32(n), err
	case uint16:
		return int32(n), err
	case uint32:
		if n > math.MaxInt32 {
			err = OverflowError{ToType: "int32", Value: v}
		}
		return int32(n), err
	case uint64:
		if n > math.MaxInt32 {
			err = OverflowError{ToType: "int32", Value: v}
		}
		return int32(n), err
	case uint:
		if n > math.MaxInt32 {
			err = OverflowError{ToType: "int32", Value: v}
		}
		return int32(n), err
	}

	return 0, InvalidTypeError{ToType: "int32", Value: v}
}
