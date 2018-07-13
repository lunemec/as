// +build 386 arm

package as

import (
	"fmt"
	"math"
)

// Int64 converts any numeric type to int64, and returns
// error if there was overflow.
func Int64(v interface{}) (int64, error) {
	var err error
	var errMsg = "%d (%T) overflows int64"

	switch n := v.(type) {
	case int8:
		return int64(n), err
	case int16:
		return int64(n), err
	case int32:
		return int64(n), err
	case int64:
		return n, err
	case int:
		return int64(n), err
	case uint8:
		return int64(n), err
	case uint16:
		return int64(n), err
	case uint32:
		return int64(n), err
	case uint64:
		if n > math.MaxInt64 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return int64(n), err
	case uint:
		return int64(n), err
	}

	return 0, fmt.Errorf("invalid type %T", v)
}
