// +build 386 arm

package as_test

import (
	"math"
	"testing"

	"as"
)

func TestUint(t *testing.T) {
	assertError(t, as.Uint, int8(math.MinInt8))
	assertNoError(t, as.Uint, int8(math.MaxInt8))
	pointerToMaxInt8 := int8(math.MaxInt8)
	assertNoError(t, as.Int, &pointerToMaxInt8)

	assertError(t, as.Uint, int16(math.MinInt16))
	assertNoError(t, as.Uint, int16(math.MaxInt16))

	assertError(t, as.Uint, int32(math.MinInt32))
	assertNoError(t, as.Uint, int32(math.MaxInt32))

	assertError(t, as.Uint, int64(math.MinInt64))
	assertError(t, as.Uint, int64(math.MaxInt64))

	assertError(t, as.Uint, int(math.MinInt32))
	assertNoError(t, as.Uint, int(math.MaxInt32))

	assertNoError(t, as.Uint, uint8(0))
	assertNoError(t, as.Uint, uint8(math.MaxUint8))

	assertNoError(t, as.Uint, uint16(0))
	assertNoError(t, as.Uint, uint16(math.MaxUint16))

	assertNoError(t, as.Uint, uint32(0))
	assertNoError(t, as.Uint, uint32(math.MaxUint32))

	assertNoError(t, as.Uint, uint64(0))
	assertError(t, as.Uint, uint64(math.MaxUint64))

	assertNoError(t, as.Uint, uint(0))
	assertNoError(t, as.Uint, uint(math.MaxUint32))
}
