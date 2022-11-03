package as

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

// SliceT is generic function to convert between slices
// of numeric types while checking for overflow on each item.
// NOTE: this function is quite slow, and especially when high %
// of numbers overflow, the allocations are significant.
func SliceT[From Number, To Number](src []From) ([]To, error) {
	var errOut error

	dst := make([]To, len(src))
	for i, v := range src {
		convertedValue, err := T[To](v)
		if err != nil {
			errOut = multierror.Append(errOut, errors.Wrapf(err, "at index [%d]", i))
			// We can skip the value, the slice will contain default for type To.
			continue
		}

		dst[i] = convertedValue
	}
	return dst, errOut
}

// SliceNumber is type constraint for SliceTUnchecked
type SliceNumber interface {
	constraints.Integer | constraints.Float
}

// SliceTUnchecked is generic function to convert between slices
// of any numeric type without overflow check.
func SliceTUnchecked[From SliceNumber, To SliceNumber](src []From) []To {
	dst := make([]To, len(src))
	for i, v := range src {
		dst[i] = To(v)
	}
	return dst
}
