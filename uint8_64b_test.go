// +build amd64 arm64

package as_test

import (
	"math"
	"testing"

	"github.com/lunemec/as"
)

func TestUint8(t *testing.T) {
	assertError(t, as.Uint8, int8(math.MinInt8))
	assertNoError(t, as.Uint8, int8(math.MaxInt8))
	pointerToMaxInt8 := int8(math.MaxInt8)
	assertNoError(t, as.Int, &pointerToMaxInt8)

	assertError(t, as.Uint8, int16(math.MinInt16))
	assertError(t, as.Uint8, int16(math.MaxInt16))

	assertError(t, as.Uint8, int32(math.MinInt32))
	assertError(t, as.Uint8, int32(math.MaxInt32))

	assertError(t, as.Uint8, int64(math.MinInt64))
	assertError(t, as.Uint8, int64(math.MaxInt64))

	assertError(t, as.Uint8, int(math.MinInt64))
	assertError(t, as.Uint8, int(math.MaxInt64))

	assertNoError(t, as.Uint8, uint8(0))
	assertNoError(t, as.Uint8, uint8(math.MaxUint8))

	assertNoError(t, as.Uint8, uint16(0))
	assertError(t, as.Uint8, uint16(math.MaxUint16))

	assertNoError(t, as.Uint8, uint32(0))
	assertError(t, as.Uint8, uint32(math.MaxUint32))

	assertNoError(t, as.Uint8, uint64(0))
	assertError(t, as.Uint8, uint64(math.MaxUint64))

	assertNoError(t, as.Uint8, uint(0))
	assertError(t, as.Uint8, uint(math.MaxUint64))
}
