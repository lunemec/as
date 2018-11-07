// +build amd64 arm64

package as

import (
	"math"
)

// Int converts any numeric type to int, and returns
// error if there was overflow.
func Int(v interface{}) (int, error) {
	var err error
	v = indirect(v)

	switch n := v.(type) {
	case int8:
		return int(n), err
	case int16:
		return int(n), err
	case int32:
		return int(n), err
	case int64:
		return int(n), err
	case int:
		return int(n), err
	case uint8:
		return int(n), err
	case uint16:
		return int(n), err
	case uint32:
		return int(n), err
	case uint64:
		if n > math.MaxInt64 {
			err = OverflowError{ToType: "int", Value: v}
		}
		return int(n), err
	case uint:
		if n > math.MaxInt64 {
			err = OverflowError{ToType: "int", Value: v}
		}
		return int(n), err
	}

	return 0, InvalidTypeError{ToType: "int", Value: v}
}
