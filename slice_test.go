package as_test

import (
	"io"
	"math"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/lunemec/as/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceT(t *testing.T) {
	tests := []struct {
		name     string
		src      []uint64
		expected []int32
		wantErr  bool
	}{
		{
			name:     "exact boundaries",
			src:      []uint64{0, 1, 2, math.MaxInt32},
			expected: []int32{0, 1, 2, math.MaxInt32},
		},
		{
			name:     "all fitted positions are converted",
			src:      []uint64{0, 1, 2, math.MaxInt32, 1, 2, 3},
			expected: []int32{0, 1, 2, math.MaxInt32, 1, 2, 3},
		},
		{
			name:     "overflow is zeroed",
			src:      []uint64{0, 1, 2, math.MaxUint32},
			expected: []int32{0, 1, 2, 0},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := as.SliceT[uint64, int32](tt.src)
			assert.Equal(t, tt.expected, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	t.Run("nil source returns initialized slice", func(t *testing.T) {
		var src []uint64
		got, err := as.SliceT[uint64, int32](src)
		require.NoError(t, err)
		assert.NotNil(t, got)
		assert.Empty(t, got)
	})

	t.Run("empty source returns initialized slice", func(t *testing.T) {
		got, err := as.SliceT[uint64, int32]([]uint64{})
		require.NoError(t, err)
		assert.NotNil(t, got)
		assert.Empty(t, got)
	})
}

func TestSliceTInto(t *testing.T) {
	tests := []struct {
		name     string
		dst      []int8
		src      []int16
		expected []int8
		wantErr  error
	}{
		{
			name:     "exact size",
			dst:      []int8{9, 9},
			src:      []int16{1, 2},
			expected: []int8{1, 2},
		},
		{
			name:     "oversized destination keeps suffix",
			dst:      []int8{9, 9, 7},
			src:      []int16{1, 2},
			expected: []int8{1, 2, 7},
		},
		{
			name:     "short destination is unchanged",
			dst:      []int8{9},
			src:      []int16{1, 2},
			expected: []int8{9},
			wantErr:  io.ErrShortBuffer,
		},
		{
			name:     "empty source leaves destination unchanged",
			dst:      []int8{7},
			src:      []int16{},
			expected: []int8{7},
		},
		{
			name: "nil source and destination",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := append([]int8(nil), tt.dst...)
			if tt.dst != nil && dst == nil {
				dst = []int8{}
			}

			err := as.SliceTInto(dst, tt.src)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.expected, dst)
		})
	}
}

func TestSliceTIntoAggregatesOverflows(t *testing.T) {
	dst := []int8{9, 9, 9, 7}
	src := []int16{math.MaxInt8, math.MaxInt8 + 1, math.MinInt8 - 1}

	err := as.SliceTInto(dst, src)

	assert.Equal(t, []int8{math.MaxInt8, 0, 0, 7}, dst)
	require.EqualError(t, err, "2 errors occurred:\n\t* at index [1]: 128 (int16) overflows int8\n\t* at index [2]: -129 (int16) overflows int8\n\n")

	var aggregate *multierror.Error
	require.ErrorAs(t, err, &aggregate)
	require.Len(t, aggregate.Errors, 2)
	for i, indexedErr := range aggregate.Errors {
		var overflow as.OverflowError
		require.ErrorAs(t, indexedErr, &overflow)
		assert.Equal(t, "int8", overflow.ToType)
		assert.EqualValues(t, src[i+1], overflow.Value)
	}
}

func TestSliceTIntoNamedAndUintptr(t *testing.T) {
	t.Run("named integers", func(t *testing.T) {
		dst := []namedInt8{9, 9}
		err := as.SliceTInto(dst, []namedInt16{1, 2})
		require.NoError(t, err)
		assert.Equal(t, []namedInt8{1, 2}, dst)
	})

	t.Run("uintptr", func(t *testing.T) {
		dst := []uint64{9, 9}
		err := as.SliceTInto(dst, []uintptr{0, 1})
		require.NoError(t, err)
		assert.Equal(t, []uint64{0, 1}, dst)
	})
}

func TestSliceTUnchecked(t *testing.T) {
	tests := []struct {
		name     string
		src      []uint64
		expected []int32
	}{
		{
			name:     "values fit",
			src:      []uint64{0, 1, 2, math.MaxInt32},
			expected: []int32{0, 1, 2, math.MaxInt32},
		},
		{
			name:     "values wrap",
			src:      []uint64{0, 1, 2, math.MaxUint32},
			expected: []int32{0, 1, 2, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, as.SliceTUnchecked[uint64, int32](tt.src))
		})
	}

	t.Run("uint8 fit and wrap", func(t *testing.T) {
		assert.Equal(t, []uint8{0, 1, 2, math.MaxUint8}, as.SliceTUnchecked[uint64, uint8]([]uint64{0, 1, 2, math.MaxUint8}))
		assert.Equal(t, []uint8{0, 1, 2, math.MaxUint8}, as.SliceTUnchecked[uint64, uint8]([]uint64{0, 1, 2, math.MaxInt32}))
	})

	t.Run("floating point", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, as.SliceTUnchecked[float64, int]([]float64{1, 2}))
	})

	t.Run("nil source returns initialized slice", func(t *testing.T) {
		var src []uint64
		got := as.SliceTUnchecked[uint64, int32](src)
		assert.NotNil(t, got)
		assert.Empty(t, got)
	})

	t.Run("empty source returns initialized slice", func(t *testing.T) {
		got := as.SliceTUnchecked[uint64, int32]([]uint64{})
		assert.NotNil(t, got)
		assert.Empty(t, got)
	})
}

func TestSliceTUncheckedInto(t *testing.T) {
	tests := []struct {
		name     string
		dst      []int32
		src      []uint64
		expected []int32
		wantErr  error
	}{
		{
			name:     "exact size",
			dst:      []int32{9, 9},
			src:      []uint64{1, 2},
			expected: []int32{1, 2},
		},
		{
			name:     "values wrap",
			dst:      []int32{9},
			src:      []uint64{math.MaxUint32},
			expected: []int32{-1},
		},
		{
			name:     "oversized destination keeps suffix",
			dst:      []int32{9, 9, 7},
			src:      []uint64{1, 2},
			expected: []int32{1, 2, 7},
		},
		{
			name:     "short destination is unchanged",
			dst:      []int32{9},
			src:      []uint64{1, 2},
			expected: []int32{9},
			wantErr:  io.ErrShortBuffer,
		},
		{
			name: "nil source and destination",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := append([]int32(nil), tt.dst...)
			if tt.dst != nil && dst == nil {
				dst = []int32{}
			}

			err := as.SliceTUncheckedInto(dst, tt.src)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.expected, dst)
		})
	}

	t.Run("floating point", func(t *testing.T) {
		dst := []int{9, 9}
		err := as.SliceTUncheckedInto(dst, []float64{1, 2})
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2}, dst)
	})

	t.Run("named integers and uintptr", func(t *testing.T) {
		dst := []namedUint16{9, 9}
		err := as.SliceTUncheckedInto(dst, []namedUintptr{0, 1})
		require.NoError(t, err)
		assert.Equal(t, []namedUint16{0, 1}, dst)
	})
}

func TestSliceTIntoSuccessDoesNotAllocate(t *testing.T) {
	src := make([]int16, 128)
	for i := range src {
		src[i] = int16(i)
	}
	dst := make([]int8, len(src))

	allocations := testing.AllocsPerRun(1000, func() {
		if err := as.SliceTInto(dst, src); err != nil {
			t.Fatal(err)
		}
	})
	assert.Zero(t, allocations)
}

func TestSliceTUncheckedIntoDoesNotAllocate(t *testing.T) {
	src := []uint64{0, 1, 2, math.MaxUint32}
	dst := make([]int, len(src))

	allocations := testing.AllocsPerRun(1000, func() {
		if err := as.SliceTUncheckedInto(dst, src); err != nil {
			t.Fatal(err)
		}
	})
	assert.Zero(t, allocations)
}

func TestAllocatingSliceConversionsAllocateOnce(t *testing.T) {
	src := []uint64{0, 1, 2, 3}
	var (
		checked   []int
		unchecked []int
		err       error
	)

	checkedAllocations := testing.AllocsPerRun(1000, func() {
		checked, err = as.SliceT[uint64, int](src)
	})
	require.NoError(t, err)
	assert.Len(t, checked, len(src))
	assert.Equal(t, float64(1), checkedAllocations)

	uncheckedAllocations := testing.AllocsPerRun(1000, func() {
		unchecked = as.SliceTUnchecked[uint64, int](src)
	})
	assert.Len(t, unchecked, len(src))
	assert.Equal(t, float64(1), uncheckedAllocations)
}

var (
	outSlice  []int
	outSlice8 []int8
	outErr    error
)

func BenchmarkSliceT(b *testing.B) {
	testdata := make([]uint64, 10000)
	for i := range testdata {
		testdata[i] = uint64(i)
	}
	b.ReportAllocs()
	for b.Loop() {
		outSlice, outErr = as.SliceT[uint64, int](testdata)
	}
}

func BenchmarkSliceTInto(b *testing.B) {
	testdata := make([]uint64, 10000)
	for i := range testdata {
		testdata[i] = uint64(i)
	}
	dst := make([]int, len(testdata))
	b.ReportAllocs()
	for b.Loop() {
		outErr = as.SliceTInto(dst, testdata)
	}
}

func BenchmarkSliceTOneOverflow(b *testing.B) {
	testdata := make([]uint64, 10000)
	testdata[len(testdata)-1] = math.MaxUint64
	b.ReportAllocs()
	for b.Loop() {
		outSlice, outErr = as.SliceT[uint64, int](testdata)
	}
}

func BenchmarkSliceTErrors(b *testing.B) {
	testdata := make([]uint64, 10000)
	for i := range testdata {
		testdata[i] = uint64(i)
	}
	b.ReportAllocs()
	for b.Loop() {
		outSlice8, outErr = as.SliceT[uint64, int8](testdata)
	}
}

func BenchmarkSliceTUnchecked(b *testing.B) {
	testdata := make([]uint64, 10000)
	for i := range testdata {
		testdata[i] = uint64(i)
	}
	b.ReportAllocs()
	for b.Loop() {
		outSlice = as.SliceTUnchecked[uint64, int](testdata)
	}
}

func BenchmarkSliceTUncheckedInto(b *testing.B) {
	testdata := make([]uint64, 10000)
	for i := range testdata {
		testdata[i] = uint64(i)
	}
	dst := make([]int, len(testdata))
	b.ReportAllocs()
	for b.Loop() {
		outErr = as.SliceTUncheckedInto(dst, testdata)
	}
}
