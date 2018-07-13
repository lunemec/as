package as

import (
	"fmt"
	"math"
)

// Int32 converts any numeric type to int32, and returns
// error if there was overflow.
func Int32(v interface{}) (int32, error) {
	var err error
	var errMsg = "%d (%T) overflows int32"

	switch n := v.(type) {
	case int8:
		return int32(n), err
	case int16:
		return int32(n), err
	case int32:
		return n, err
	case int64:
		if n < math.MinInt32 || n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int32(n), err
	case int:
		if n < math.MinInt32 || n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int32(n), err
	case uint8:
		return int32(n), err
	case uint16:
		return int32(n), err
	case uint32:
		if n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int32(n), err
	case uint64:
		if n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int32(n), err
	case uint:
		if n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int32(n), err
	}

	return 0, fmt.Errorf("invalid type %T", v)
}
