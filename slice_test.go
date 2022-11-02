package as_test

import (
	"math"
	"testing"

	"github.com/lunemec/as"
	"github.com/stretchr/testify/assert"
)

func TestSliceT(t *testing.T) {
	assertSliceTNoError(t, as.SliceT[uint64, int32], []uint64{0, 1, 2, math.MaxInt32}, []int32{0, 1, 2, 2147483647})
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
