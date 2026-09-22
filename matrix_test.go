package as_test

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/lunemec/as"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegerConversionMatrix(t *testing.T) {
	runIntegerMatrix[int](t, "int")
	runIntegerMatrix[int8](t, "int8")
	runIntegerMatrix[int16](t, "int16")
	runIntegerMatrix[int32](t, "int32")
	runIntegerMatrix[int64](t, "int64")
	runIntegerMatrix[uint](t, "uint")
	runIntegerMatrix[uint8](t, "uint8")
	runIntegerMatrix[uint16](t, "uint16")
	runIntegerMatrix[uint32](t, "uint32")
	runIntegerMatrix[uint64](t, "uint64")
	runIntegerMatrix[uintptr](t, "uintptr")
}

func runIntegerMatrix[To as.Number](t *testing.T, name string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		testIntegerSources[To](t)
	})
}

func testIntegerSources[To as.Number](t *testing.T) {
	t.Helper()

	testMatrixValue[To](t, int8(math.MinInt8))
	testMatrixValue[To](t, int8(math.MaxInt8))
	testMatrixValue[To](t, int16(math.MinInt16))
	testMatrixValue[To](t, int16(math.MaxInt16))
	testMatrixValue[To](t, int32(math.MinInt32))
	testMatrixValue[To](t, int32(math.MaxInt32))
	testMatrixValue[To](t, int64(math.MinInt64))
	testMatrixValue[To](t, int64(math.MaxInt64))
	testMatrixValue[To](t, int(math.MinInt))
	testMatrixValue[To](t, int(math.MaxInt))

	testMatrixValue[To](t, uint8(0))
	testMatrixValue[To](t, uint8(math.MaxUint8))
	testMatrixValue[To](t, uint16(0))
	testMatrixValue[To](t, uint16(math.MaxUint16))
	testMatrixValue[To](t, uint32(0))
	testMatrixValue[To](t, uint32(math.MaxUint32))
	testMatrixValue[To](t, uint64(0))
	testMatrixValue[To](t, uint64(math.MaxUint32))
	testMatrixValue[To](t, uint64(math.MaxUint64))
	testMatrixValue[To](t, uint(0))
	if ^uint(0) > uint(math.MaxUint32) {
		testMatrixValue[To](t, uint(math.MaxUint32))
	}
	testMatrixValue[To](t, ^uint(0))
	testMatrixValue[To](t, uintptr(0))
	testMatrixValue[To](t, ^uintptr(0))
}

func FuzzIntegerConversions(f *testing.F) {
	f.Add(int64(0), uint64(0))
	f.Add(int64(-1), uint64(1))
	f.Add(int64(math.MinInt8-1), uint64(math.MaxUint8+1))
	f.Add(int64(math.MinInt16-1), uint64(math.MaxUint16+1))
	f.Add(int64(math.MinInt32-1), uint64(math.MaxUint32+1))
	f.Add(int64(math.MinInt64), uint64(math.MaxUint64))
	f.Add(int64(math.MaxInt64), uint64(math.MaxInt64))

	f.Fuzz(func(t *testing.T, signed int64, unsigned uint64) {
		fuzzIntegerTargets(t, int(signed))
		fuzzIntegerTargets(t, int8(signed))
		fuzzIntegerTargets(t, int16(signed))
		fuzzIntegerTargets(t, int32(signed))
		fuzzIntegerTargets(t, signed)
		fuzzIntegerTargets(t, uint(unsigned))
		fuzzIntegerTargets(t, uint8(unsigned))
		fuzzIntegerTargets(t, uint16(unsigned))
		fuzzIntegerTargets(t, uint32(unsigned))
		fuzzIntegerTargets(t, unsigned)
		fuzzIntegerTargets(t, uintptr(unsigned))
	})
}

func fuzzIntegerTargets[From as.Number](t *testing.T, value From) {
	t.Helper()
	fuzzIntegerTarget[int](t, value)
	fuzzIntegerTarget[int8](t, value)
	fuzzIntegerTarget[int16](t, value)
	fuzzIntegerTarget[int32](t, value)
	fuzzIntegerTarget[int64](t, value)
	fuzzIntegerTarget[uint](t, value)
	fuzzIntegerTarget[uint8](t, value)
	fuzzIntegerTarget[uint16](t, value)
	fuzzIntegerTarget[uint32](t, value)
	fuzzIntegerTarget[uint64](t, value)
	fuzzIntegerTarget[uintptr](t, value)
}

func fuzzIntegerTarget[To, From as.Number](t *testing.T, value From) {
	t.Helper()
	got, err := as.T[To](value)
	if got != To(value) {
		t.Fatalf("T[%T](%v) = %v, want wrapped value %v", got, value, got, To(value))
	}
	if fits := integerFits[To](value); fits != (err == nil) {
		t.Fatalf("T[%T](%T(%v)) error = %v, fits = %v", got, value, value, err, fits)
	}
}

func testMatrixValue[To, From as.Number](t *testing.T, value From) {
	t.Helper()

	t.Run(fmt.Sprintf("%T(%v)", value, value), func(t *testing.T) {
		got, err := as.T[To](value)
		assert.Equal(t, To(value), got)

		if integerFits[To](value) {
			require.NoError(t, err)
			return
		}

		require.Error(t, err)
		var overflow as.OverflowError
		require.ErrorAs(t, err, &overflow)
		assert.Equal(t, reflect.TypeFor[To]().String(), overflow.ToType)
		assert.Equal(t, value, overflow.Value)
	})
}

func integerFits[To, From as.Number](value From) bool {
	number := bigInteger(value)
	bits := reflect.TypeFor[To]().Bits()
	limit := new(big.Int).Lsh(big.NewInt(1), uint(bits))
	minimum := big.NewInt(0)

	switch reflect.TypeFor[To]().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		limit.Rsh(limit, 1)
		minimum.Neg(new(big.Int).Set(limit))
	}

	maximum := new(big.Int).Sub(limit, big.NewInt(1))
	return number.Cmp(minimum) >= 0 && number.Cmp(maximum) <= 0
}

func bigInteger[T as.Number](value T) *big.Int {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return big.NewInt(reflected.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return new(big.Int).SetUint64(reflected.Uint())
	default:
		panic("unreachable integer kind")
	}
}
