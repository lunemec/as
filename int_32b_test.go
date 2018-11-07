// +build 386 arm

package as_test

import (
	"math"
	"testing"

	"as"
)

func TestInt(t *testing.T) {
	assertNoError(t, as.Int, int8(math.MinInt8))
	assertNoError(t, as.Int, int8(math.MaxInt8))
	pointerToMaxInt8 := int8(math.MaxInt8)
	assertNoError(t, as.Int, &pointerToMaxInt8)

	assertNoError(t, as.Int, int16(math.MinInt16))
	assertNoError(t, as.Int, int16(math.MaxInt16))

	assertNoError(t, as.Int, int32(math.MinInt32))
	assertNoError(t, as.Int, int32(math.MaxInt32))

	assertError(t, as.Int, int64(math.MinInt64))
	assertError(t, as.Int, int64(math.MaxInt64))

	assertNoError(t, as.Int, int(math.MinInt32))
	assertNoError(t, as.Int, int(math.MaxInt32))

	assertNoError(t, as.Int, uint8(0))
	assertNoError(t, as.Int, uint8(math.MaxUint8))

	assertNoError(t, as.Int, uint16(0))
	assertNoError(t, as.Int, uint16(math.MaxUint16))

	assertNoError(t, as.Int, uint32(0))
	assertError(t, as.Int, uint32(math.MaxUint32))

	assertNoError(t, as.Int, uint64(0))
	assertError(t, as.Int, uint64(math.MaxUint32))
	assertError(t, as.Int, uint64(math.MaxUint64))

	assertNoError(t, as.Int, uint(0))
	assertError(t, as.Int, uint(math.MaxUint32))
	assertError(t, as.Int, uint(math.MaxUint32))
}
