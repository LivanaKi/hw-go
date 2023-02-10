package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtoi(t *testing.T) {
	const str, want = "43", 43
	got, err := strconv.Atoi(str)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.True(t, err == nil)
	require.Nil(t, err)
}
