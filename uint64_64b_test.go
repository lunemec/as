// +build amd64 arm64

package as_test

import (
	"math"
	"testing"

	"as"
)

func TestUint64(t *testing.T) {
	assertError(t, as.Uint64, int8(math.MinInt8))
	assertNoError(t, as.Uint64, int8(math.MaxInt8))
	pointerToMaxInt8 := int8(math.MaxInt8)
	assertNoError(t, as.Int, &pointerToMaxInt8)

	assertError(t, as.Uint64, int16(math.MinInt16))
	assertNoError(t, as.Uint64, int16(math.MaxInt16))

	assertError(t, as.Uint64, int32(math.MinInt32))
	assertNoError(t, as.Uint64, int32(math.MaxInt32))

	assertError(t, as.Uint64, int64(math.MinInt64))
	assertNoError(t, as.Uint64, int64(math.MaxInt64))

	assertError(t, as.Uint64, int(math.MinInt64))
	assertNoError(t, as.Uint64, int(math.MaxInt64))

	assertNoError(t, as.Uint64, uint8(0))
	assertNoError(t, as.Uint64, uint8(math.MaxUint8))

	assertNoError(t, as.Uint64, uint16(0))
	assertNoError(t, as.Uint64, uint16(math.MaxUint16))

	assertNoError(t, as.Uint64, uint32(0))
	assertNoError(t, as.Uint64, uint32(math.MaxUint32))

	assertNoError(t, as.Uint64, uint64(0))
	assertNoError(t, as.Uint64, uint64(math.MaxUint64))

	assertNoError(t, as.Uint64, uint(0))
	assertNoError(t, as.Uint64, uint(math.MaxUint64))
}
