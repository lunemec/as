// +build amd64 arm64

package as_test

import (
	"math"
	"testing"

	"as"
)

func TestInt16(t *testing.T) {
	assertNoError(t, as.Int16, int8(math.MinInt8))
	assertNoError(t, as.Int16, int8(math.MaxInt8))

	assertNoError(t, as.Int16, int16(math.MinInt16))
	assertNoError(t, as.Int16, int16(math.MaxInt16))

	assertError(t, as.Int16, int32(math.MinInt32))
	assertError(t, as.Int16, int32(math.MaxInt32))

	assertError(t, as.Int16, int64(math.MinInt64))
	assertError(t, as.Int16, int64(math.MaxInt64))

	assertError(t, as.Int16, int(math.MinInt64))
	assertError(t, as.Int16, int(math.MaxInt64))

	assertNoError(t, as.Int16, uint8(0))
	assertNoError(t, as.Int16, uint8(math.MaxUint8))

	assertNoError(t, as.Int16, uint16(0))
	assertError(t, as.Int16, uint16(math.MaxUint16))

	assertNoError(t, as.Int16, uint32(0))
	assertError(t, as.Int16, uint32(math.MaxUint32))

	assertNoError(t, as.Int16, uint64(0))
	assertError(t, as.Int16, uint64(math.MaxUint64))

	assertNoError(t, as.Int16, uint(0))
	assertError(t, as.Int16, uint(math.MaxUint64))
}
