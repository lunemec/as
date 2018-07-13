// +build 386 arm

package as

import (
	"fmt"
	"math"
)

// Int converts any numeric type to int, and returns
// error if there was overflow.
func Int(v interface{}) (int, error) {
	var err error
	var errMsg = "%d (%T) overflows int"

	switch n := v.(type) {
	case int8:
		return int(n), err
	case int16:
		return int(n), err
	case int32:
		return int(n), err
	case int64:
		if n < math.MinInt32 || n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int(n), err
	case int:
		return int(n), err
	case uint8:
		return int(n), err
	case uint16:
		return int(n), err
	case uint32:
		if n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int(n), err
	case uint64:
		if n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int(n), err
	case uint:
		if n > math.MaxInt32 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int(n), err
	}

	return 0, fmt.Errorf("invalid type %T", v)
}