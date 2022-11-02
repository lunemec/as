//go:build 386 || arm
// +build 386 arm

package as_test

import (
	"testing"
)

func TestT32b(t *testing.T) {
	assertNoError(t, as.T[int], int64(math.MaxInt32))
	assertNoError(t, as.T[int], int64(math.MinInt32))
	assertError(t, as.T[int], int64(math.MaxInt64))
	assertError(t, as.T[int], int64(math.MinInt64))
	assertError(t, as.T[int], uint64(math.MaxUint64))
}
