//go:build amd64 || arm64
// +build amd64 arm64

package as_test

import (
	"math"
	"testing"

	"github.com/lunemec/as"
)

func TestInt16(t *testing.T) {
	assertNoError(t, as.Int16, int8(math.MinInt8))
	assertNoError(t, as.Int16, int8(math.MaxInt8))
	pointerToMaxInt8 := int8(math.MaxInt8)
	assertNoError(t, as.Int, &pointerToMaxInt8)

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

var out16 int16

// BenchmarkAs16-8   	36798218	        32.61 ns/op	      39 B/op	       1 allocs/op

func BenchmarkAs16(b *testing.B) {
	var t int16
	for n := 0; n < b.N; n++ {
		t, _ = as.Int16(n)
	}

	out16 = t
}

// BenchmarkInt16-8   	1000000000	         0.3211 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInt16(b *testing.B) {
	var t int16
	for n := 0; n < b.N; n++ {
		t = int16(n)
	}

	out16 = t
}
