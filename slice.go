package as

import (
	"github.com/hashicorp/go-multierror"
	"golang.org/x/exp/constraints"
)

type SliceNumber interface {
	constraints.Integer | constraints.Float
}

func SliceT[From Number, To Number](src []From) ([]To, error) {
	var errOut error

	dst := make([]To, len(src))
	for i, v := range src {
		convertedValue, err := T[To](v)
		if err != nil {
			errOut = multierror.Append(errOut, err)
			// We can skip the value, the slice will contain default for type To.
			continue
		}

		dst[i] = convertedValue
	}
	return dst, errOut
}

func SliceTUnchecked[From SliceNumber, To SliceNumber](src []From) []To {
	dst := make([]To, len(src))
	for i, v := range src {
		dst[i] = To(v)
	}
	return dst
}
