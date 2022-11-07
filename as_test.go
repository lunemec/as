package as_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/lunemec/as"
)

// assertError tests if function called with given parameters raises error, if not
// it is considered failed test.
func assertError(t *testing.T, fn interface{}, args interface{}) {
	t.Helper()
	outType, err := callfn(fn, args)
	if err == nil {
		t.Errorf("There should be error when converting %+v (%T) to %v but err is %v", args, args, outType, err)
	}
}

// assertNoError tests if function called with given parameters raises error, if not
// it is considered failed test.
func assertNoError(t *testing.T, fn interface{}, args interface{}) {
	t.Helper()
	outType, err := callfn(fn, args)
	if err != nil {
		t.Errorf("There should be NO error when converting %+v (%T) to %v: %s", args, args, outType, err)
	}
}

func callfn(fn interface{}, args interface{}) (reflect.Type, interface{}) {
	v := reflect.ValueOf(fn)
	if t := v.Kind(); t != reflect.Func {
		panic("assertError expect function as 2nd argument!")
	}
	outValue := v.Call([]reflect.Value{reflect.ValueOf(args)})
	outType := outValue[0].Type()
	err := outValue[1].Interface()
	return outType, err
}

func TestT(t *testing.T) {
	assertNoError(t, as.T[int], int64(math.MaxInt64))
	assertNoError(t, as.T[int], int64(math.MinInt64))
	assertError(t, as.T[int], uint64(math.MaxUint64))
	// uintptr is not supported, will error
	assertError(t, as.T[uintptr], uint64(math.MaxUint64))
}

func TestTCustomType(t *testing.T) {
	type aliasedInt int

	assertNoError(t, as.T[int], aliasedInt(math.MaxInt64))
	assertNoError(t, as.T[int], aliasedInt(math.MinInt64))
}

var out int

// BenchmarkT-8   	91144112	        12.90 ns/op	       8 B/op	       0 allocs/op
// After changes to indirect:
// BenchmarkT-8   	40429498	        29.37 ns/op	      16 B/op	       1 allocs/op
func BenchmarkT(b *testing.B) {
	var t int
	for n := 0; n < b.N; n++ {
		t, _ = as.T[int](n)
	}

	out = t
}

// BenchmarkInt-8   	108703077	        10.98 ns/op	       8 B/op	       0 allocs/op
// After changes to indirect:
// BenchmarkInt-8   	42888086	        28.00 ns/op	      16 B/op	       1 allocs/op
func BenchmarkInt(b *testing.B) {
	var t int
	for n := 0; n < b.N; n++ {
		t, _ = as.Int(n)
	}

	out = t
}
