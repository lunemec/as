// +build 386 arm

package as_test

import (
	"math"
	"testing"

	"as"
)

func TestInt8(t *testing.T) {
	assertNoError(t, as.Int8, int8(math.MinInt8))
	assertNoError(t, as.Int8, int8(math.MaxInt8))

	assertError(t, as.Int8, int16(math.MinInt16))
	assertError(t, as.Int8, int16(math.MaxInt16))

	assertError(t, as.Int8, int32(math.MinInt32))
	assertError(t, as.Int8, int32(math.MaxInt32))

	assertError(t, as.Int8, int64(math.MinInt64))
	assertError(t, as.Int8, int64(math.MaxInt64))

	assertError(t, as.Int8, int(math.MinInt32))
	assertError(t, as.Int8, int(math.MaxInt32))

	assertNoError(t, as.Int8, uint8(0))
	assertError(t, as.Int8, uint8(math.MaxUint8))

	assertNoError(t, as.Int8, uint16(0))
	assertError(t, as.Int8, uint16(math.MaxUint16))

	assertNoError(t, as.Int8, uint32(0))
	assertError(t, as.Int8, uint32(math.MaxUint32))

	assertNoError(t, as.Int8, uint64(0))
	assertError(t, as.Int8, uint64(math.MaxUint64))

	assertNoError(t, as.Int8, uint(0))
	assertError(t, as.Int8, uint(math.MaxUint32))
}
