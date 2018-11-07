// +build 386 arm

package as

// Uint64 converts any numeric type to uint64, and returns
// error if there was overflow.
func Uint64(v interface{}) (uint64, error) {
	var err error
	v = indirect(v)

	switch n := v.(type) {
	case int8:
		if n < 0 {
			err = OverflowError{ToType: "uint64", Value: v}
		}
		return uint64(n), err
	case int16:
		if n < 0 {
			err = OverflowError{ToType: "uint64", Value: v}
		}
		return uint64(n), err
	case int32:
		if n < 0 {
			err = OverflowError{ToType: "uint64", Value: v}
		}
		return uint64(n), err
	case int64:
		if n < 0 {
			err = OverflowError{ToType: "uint64", Value: v}
		}
		return uint64(n), err
	case int:
		if n < 0 {
			err = OverflowError{ToType: "uint64", Value: v}
		}
		return uint64(n), err
	case uint8:
		return uint64(n), err
	case uint16:
		return uint64(n), err
	case uint32:
		return uint64(n), err
	case uint64:
		return n, err
	case uint:
		return uint64(n), err
	}

	return 0, InvalidTypeError{ToType: "uint64", Value: v}
}
