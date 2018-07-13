package as

import (
	"fmt"
	"math"
)

// Int8 converts any numeric type to int8, and returns
// error if there was overflow.
func Int8(v interface{}) (int8, error) {
	var err error
	var errMsg = "%d (%T) overflows int8"

	switch n := v.(type) {
	case int8:
		return n, err
	case int16:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case int32:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case int64:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case int:
		if n < math.MinInt8 || n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case uint8:
		if n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case uint16:
		if n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case uint32:
		if n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case uint64:
		if n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	case uint:
		if n > math.MaxInt8 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int8(n), err
	}

	return 0, fmt.Errorf("invalid type %T", v)
}
