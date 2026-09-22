package as_test

import (
	"testing"

	"github.com/lunemec/as"
	"github.com/stretchr/testify/require"
)

func assertNoError[From, To as.Number](t *testing.T, fn func(From) (To, error), value From) {
	t.Helper()
	_, err := fn(value)
	require.NoError(t, err)
}

func assertError[From, To as.Number](t *testing.T, fn func(From) (To, error), value From) {
	t.Helper()
	_, err := fn(value)
	require.Error(t, err)
}
