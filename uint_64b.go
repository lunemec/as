// +build amd64 arm64

package as

import (
	"fmt"
)

// Uint converts any numeric type to uint64, and returns
// error if there was overflow.
func Uint(v interface{}) (uint, error) {
	var err error
	var errMsg = "%d (%T) overflows uint"

	switch n := v.(type) {
	case int8:
		if n < 0 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return uint(n), err
	case int16:
		if n < 0 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return uint(n), err
	case int32:
		if n < 0 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return uint(n), err
	case int64:
		if n < 0 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return uint(n), err
	case int:
		if n < 0 {
			err = fmt.Errorf(errMsg, n, v)
		}
		return uint(n), err
	case uint8:
		return uint(n), err
	case uint16:
		return uint(n), err
	case uint32:
		return uint(n), err
	case uint64:
		return uint(n), err
	case uint:
		return n, err
	}

	return 0, fmt.Errorf("invalid type %T", v)
}
