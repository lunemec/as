package as

import (
	"math"
)

// Int16 converts any numeric type to int16, and returns
// error if there was overflow.
func Int16(v interface{}) (int16, error) {
	var err error

	switch n := v.(type) {
	case int8:
		return int16(n), err
	case int16:
		return n, err
	case int32:
		if n < math.MinInt16 || n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	case int64:
		if n < math.MinInt16 || n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	case int:
		if n < math.MinInt16 || n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	case uint8:
		return int16(n), err
	case uint16:
		if n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	case uint32:
		if n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	case uint64:
		if n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	case uint:
		if n > math.MaxInt16 {
			err = OverflowError{ToType: "int16", Value: v}
		}
		return int16(n), err
	}

	return 0, InvalidTypeError{ToType: "int16", Value: v}
}
