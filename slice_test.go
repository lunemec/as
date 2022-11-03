package as_test

import (
	"math"
	"testing"

	"github.com/lunemec/as"
	"github.com/stretchr/testify/assert"
)

func TestSliceT(t *testing.T) {
	assertSliceTNoError(t, as.SliceT[uint64, int32], []uint64{0, 1, 2, math.MaxInt32}, []int32{0, 1, 2, 2147483647})
	assertSliceTNoError(t, as.SliceT[uint64, int32], []uint64{0, 1, 2, math.MaxInt32, 1, 2, 3}, []int32{0, 1, 2, 2147483647, 1, 2, 3})
	assertSliceTError(t, as.SliceT[uint64, int32], []uint64{0, 1, 2, math.MaxUint32}, []int32{0, 1, 2, 0})
}

func TestSliceTUnchecked(t *testing.T) {
	// No overflow.
	assertSliceTUnchecked(t, as.SliceTUnchecked[uint64, int32], []uint64{0, 1, 2, math.MaxInt32}, []int32{0, 1, 2, 2147483647})
	// This overflows.
	assertSliceTUnchecked(t, as.SliceTUnchecked[uint64, int32], []uint64{0, 1, 2, math.MaxUint32}, []int32{0, 1, 2, -1})

	// No overflow.
	assertSliceTUnchecked(t, as.SliceTUnchecked[uint64, uint8], []uint64{0, 1, 2, 255}, []uint8{0, 1, 2, 255})
	// This overflows.
	assertSliceTUnchecked(t, as.SliceTUnchecked[uint64, uint8], []uint64{0, 1, 2, math.MaxInt32}, []uint8{0, 1, 2, 255})
}

func assertSliceTNoError[T any, T2 any](t *testing.T, testFn func([]T) ([]T2, error), testData []T, expectedData []T2) {
	gotData, err := testFn(testData)
	assert.EqualValues(t, expectedData, gotData)
	assert.NoError(t, err)
}

func assertSliceTError[T any, T2 any](t *testing.T, testFn func([]T) ([]T2, error), testData []T, expectedData []T2) {
	gotData, err := testFn(testData)
	assert.EqualValues(t, expectedData, gotData)
	assert.Error(t, err)
}

func assertSliceTUnchecked[T any, T2 any](t *testing.T, testFn func([]T) []T2, testData []T, expectedData []T2) {
	gotData := testFn(testData)
	assert.EqualValues(t, expectedData, gotData)
}

var outSlice []int
var outSlice8 []int8

// BenchmarkSliceT-8   	    8395	    137720 ns/op	  159872 B/op	    9745 allocs/op
func BenchmarkSliceT(b *testing.B) {
	var (
		t        []int
		testdata []uint64
	)

	for i := 0; i < 10000; i++ {
		testdata = append(testdata, uint64(i))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		t, _ = as.SliceT[uint64, int](testdata)
	}

	outSlice = t
}

// BenchmarkSliceTErrors-8   	     190	   6225532 ns/op	 4644377 B/op	   78749 allocs/op
func BenchmarkSliceTErrors(b *testing.B) {
	var (
		t        []int8
		testdata []uint64
	)

	for i := 0; i < 10000; i++ {
		testdata = append(testdata, uint64(i))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		t, _ = as.SliceT[uint64, int8](testdata)
	}

	outSlice8 = t
}

// BenchmarkSliceTUnchecked-8   	  157938	      7114 ns/op	   81920 B/op	       1 allocs/op
func BenchmarkSliceTUnchecked(b *testing.B) {
	var (
		t        []int
		testdata []uint64
	)

	for i := 0; i < 10000; i++ {
		testdata = append(testdata, uint64(i))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		t = as.SliceTUnchecked[uint64, int](testdata)
	}

	outSlice = t
}
