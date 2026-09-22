package as

import (
	"fmt"
	"io"

	"github.com/hashicorp/go-multierror"
)

// SliceT converts src to a newly allocated slice, checking every value for
// overflow. Overflowing positions contain the zero value of To, and the
// returned error aggregates every overflow with its source index.
//
// SliceT performs one allocation and always returns an initialized slice,
// including for nil or empty src. Use [SliceTInto] to reuse caller-owned
// storage.
func SliceT[From Number, To Number](src []From) ([]To, error) {
	dst := make([]To, len(src))
	return dst, SliceTInto(dst, src)
}

// SliceTInto converts src into the first len(src) elements of dst, checking
// every value for overflow. It returns [io.ErrShortBuffer] without modifying
// dst when len(dst) is smaller than len(src). Elements beyond len(src) are
// left unchanged.
//
// Overflowing positions are set to the zero value of To, and the returned
// error aggregates every overflow with its source index. Successful
// conversions do not allocate. Error construction may allocate. src and dst
// must not overlap.
func SliceTInto[From Number, To Number](dst []To, src []From) error {
	if len(dst) < len(src) {
		return io.ErrShortBuffer
	}

	var errs []error
	for i, v := range src {
		converted, ok := checkedCast[To](v)
		if !ok {
			var zero To
			dst[i] = zero
			errs = append(errs, fmt.Errorf("at index [%d]: %w", i, overflowError[To](v)))
			continue
		}

		dst[i] = converted
	}

	if len(errs) == 0 {
		return nil
	}
	return &multierror.Error{Errors: errs}
}

// SliceNumber is the set of numeric types supported by [SliceTUnchecked] and
// [SliceTUncheckedInto].
type SliceNumber interface {
	Number | ~float32 | ~float64
}

// SliceTUnchecked converts src to a newly allocated slice using ordinary Go
// numeric conversions. Values that do not fit in To wrap or truncate according
// to Go's conversion rules. It performs one allocation and always returns an
// initialized slice, including for nil or empty src. Use [SliceTUncheckedInto]
// to reuse caller-owned storage.
func SliceTUnchecked[From SliceNumber, To SliceNumber](src []From) []To {
	dst := make([]To, len(src))
	sliceTUncheckedInto(dst, src)
	return dst
}

// SliceTUncheckedInto converts src into the first len(src) elements of dst
// using ordinary Go numeric conversions. It returns [io.ErrShortBuffer]
// without modifying dst when len(dst) is smaller than len(src). Elements
// beyond len(src) are left unchanged. Successful conversions do not allocate.
// src and dst must not overlap.
func SliceTUncheckedInto[From SliceNumber, To SliceNumber](dst []To, src []From) error {
	if len(dst) < len(src) {
		return io.ErrShortBuffer
	}

	sliceTUncheckedInto(dst, src)
	return nil
}

func sliceTUncheckedInto[From SliceNumber, To SliceNumber](dst []To, src []From) {
	for i, v := range src {
		dst[i] = To(v)
	}
}
