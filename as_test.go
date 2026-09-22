package as_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/lunemec/as"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type namedInt16 int16
type namedInt8 int8
type namedUint16 uint16
type namedUint8 uint8
type namedUintptr uintptr

func TestTBoundaries(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		requireCast(t, int16(math.MinInt8), int8(math.MinInt8))
		requireCast(t, int16(math.MaxInt8), int8(math.MaxInt8))
		requireOverflow(t, int16(math.MinInt8-1), int8(math.MaxInt8))
		requireOverflow(t, int16(math.MaxInt8+1), int8(math.MinInt8))
	})

	t.Run("int16", func(t *testing.T) {
		requireCast(t, int32(math.MinInt16), int16(math.MinInt16))
		requireCast(t, int32(math.MaxInt16), int16(math.MaxInt16))
		requireOverflow(t, int32(math.MinInt16-1), int16(math.MaxInt16))
		requireOverflow(t, int32(math.MaxInt16+1), int16(math.MinInt16))
	})

	t.Run("int32", func(t *testing.T) {
		requireCast(t, int64(math.MinInt32), int32(math.MinInt32))
		requireCast(t, int64(math.MaxInt32), int32(math.MaxInt32))
		requireOverflow(t, int64(math.MinInt32-1), int32(math.MaxInt32))
		requireOverflow(t, int64(math.MaxInt32+1), int32(math.MinInt32))
	})

	t.Run("int64", func(t *testing.T) {
		requireCast(t, int64(math.MinInt64), int64(math.MinInt64))
		requireCast(t, uint64(math.MaxInt64), int64(math.MaxInt64))
		requireOverflow(t, uint64(math.MaxInt64)+1, int64(math.MinInt64))
	})

	t.Run("int", func(t *testing.T) {
		requireCast(t, int64(math.MinInt), int(math.MinInt))
		requireCast(t, uint64(math.MaxInt), int(math.MaxInt))
		requireOverflow(t, uint64(math.MaxInt)+1, int(math.MinInt))
	})

	t.Run("uint8", func(t *testing.T) {
		requireCast(t, int8(0), uint8(0))
		requireCast(t, uint16(math.MaxUint8), uint8(math.MaxUint8))
		requireOverflow(t, int8(-1), uint8(math.MaxUint8))
		requireOverflow(t, uint16(math.MaxUint8)+1, uint8(0))
	})

	t.Run("uint16", func(t *testing.T) {
		requireCast(t, int8(0), uint16(0))
		requireCast(t, uint32(math.MaxUint16), uint16(math.MaxUint16))
		requireOverflow(t, int16(-1), uint16(math.MaxUint16))
		requireOverflow(t, uint32(math.MaxUint16)+1, uint16(0))
	})

	t.Run("uint32", func(t *testing.T) {
		requireCast(t, int8(0), uint32(0))
		requireCast(t, uint64(math.MaxUint32), uint32(math.MaxUint32))
		requireOverflow(t, int32(-1), uint32(math.MaxUint32))
		requireOverflow(t, uint64(math.MaxUint32)+1, uint32(0))
	})

	t.Run("uint64", func(t *testing.T) {
		requireCast(t, int8(0), uint64(0))
		requireCast(t, uint64(math.MaxUint64), uint64(math.MaxUint64))
		requireOverflow(t, int64(-1), uint64(math.MaxUint64))
	})

	t.Run("uint", func(t *testing.T) {
		max := ^uint(0)
		requireCast(t, int8(0), uint(0))
		requireCast(t, max, max)
		requireOverflow(t, int(-1), max)
		if uint64(max) < math.MaxUint64 {
			requireOverflow(t, uint64(max)+1, uint(0))
		}
	})

	t.Run("uintptr", func(t *testing.T) {
		max := ^uintptr(0)
		requireCast(t, int8(0), uintptr(0))
		requireCast(t, max, max)
		requireOverflow(t, int(-1), max)
		if uint64(max) < math.MaxUint64 {
			requireOverflow(t, uint64(max)+1, uintptr(0))
		}
	})
}

func TestTCrossSignedModuloValues(t *testing.T) {
	requireOverflow(t, int8(-1), uint16(math.MaxUint16))
	requireOverflow(t, uint8(math.MaxUint8), int8(-1))
}

func TestTNamedTypes(t *testing.T) {
	requireCast(t, namedInt16(math.MaxInt8), namedInt8(math.MaxInt8))
	requireOverflow(t, namedInt16(math.MaxInt8+1), namedInt8(math.MinInt8))
	requireCast(t, namedUint16(math.MaxUint8), namedUint8(math.MaxUint8))
	requireOverflow(t, namedUint16(math.MaxUint8+1), namedUint8(0))
	requireCast(t, namedUintptr(1), namedUintptr(1))
}

func TestOverflowError(t *testing.T) {
	got, err := as.T[int8](int16(math.MaxInt8 + 1))
	require.Error(t, err)

	var overflow as.OverflowError
	require.ErrorAs(t, err, &overflow)
	assert.Equal(t, int8(math.MinInt8), got)
	assert.Equal(t, "int8", overflow.ToType)
	assert.Equal(t, int16(math.MaxInt8+1), overflow.Value)
	assert.Equal(t, "128 (int16) overflows int8", overflow.Error())
}

func TestNamedHelpers(t *testing.T) {
	_, err := as.Int(int8(1))
	require.NoError(t, err)
	_, err = as.Int8(int16(1))
	require.NoError(t, err)
	_, err = as.Int16(int32(1))
	require.NoError(t, err)
	_, err = as.Int32(int64(1))
	require.NoError(t, err)
	_, err = as.Int64(uint32(1))
	require.NoError(t, err)
	_, err = as.Uint(uint8(1))
	require.NoError(t, err)
	_, err = as.Uint8(uint16(1))
	require.NoError(t, err)
	_, err = as.Uint16(uint32(1))
	require.NoError(t, err)
	_, err = as.Uint32(uint64(1))
	require.NoError(t, err)
	_, err = as.Uint64(uintptr(1))
	require.NoError(t, err)
	_, err = as.Uintptr(uint64(1))
	require.NoError(t, err)
}

func TestTSuccessDoesNotAllocate(t *testing.T) {
	var (
		got int8
		err error
	)
	allocations := testing.AllocsPerRun(1000, func() {
		got, err = as.T[int8](int16(math.MaxInt8))
	})
	require.NoError(t, err)
	assert.Equal(t, int8(math.MaxInt8), got)
	assert.Zero(t, allocations)
}

func FuzzInt(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxInt))
	f.Add(^uint(0))

	f.Fuzz(func(t *testing.T, value uint) {
		got, err := as.Int(value)
		if value > uint(math.MaxInt) {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, int(value), got)
	})
}

func requireCast[To, From as.Number](t *testing.T, value From, expected To) {
	t.Helper()
	got, err := as.T[To](value)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func requireOverflow[To, From as.Number](t *testing.T, value From, expected To) {
	t.Helper()
	got, err := as.T[To](value)
	require.Error(t, err)
	assert.Equal(t, expected, got)

	var overflow as.OverflowError
	require.ErrorAs(t, err, &overflow)
	assert.Equal(t, reflect.TypeFor[To]().String(), overflow.ToType)
	assert.Equal(t, value, overflow.Value)
}

var (
	benchmarkInt      int
	benchmarkInt8     int8
	benchmarkNamedInt namedInt8
	benchmarkErr      error
)

func BenchmarkT(b *testing.B) {
	value := int(123456234)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkInt, benchmarkErr = as.T[int](value)
	}
}

func BenchmarkInt(b *testing.B) {
	value := int(123456234)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkInt, benchmarkErr = as.Int(value)
	}
}

func BenchmarkUint8toInt(b *testing.B) {
	value := uint8(234)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkInt, benchmarkErr = as.Int(value)
	}
}

func BenchmarkUint64toInt(b *testing.B) {
	value := uint64(math.MaxInt)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkInt, benchmarkErr = as.Int(value)
	}
}

func BenchmarkUint64toIntOver(b *testing.B) {
	value := uint64(math.MaxUint64)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkInt, benchmarkErr = as.Int(value)
	}
}

func BenchmarkTInt16ToInt8(b *testing.B) {
	value := int16(math.MaxInt8)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkInt8, benchmarkErr = as.T[int8](value)
	}
}

func BenchmarkTNamedTypes(b *testing.B) {
	value := namedInt16(math.MaxInt8)
	b.ReportAllocs()
	for b.Loop() {
		benchmarkNamedInt, benchmarkErr = as.T[namedInt8](value)
	}
}
