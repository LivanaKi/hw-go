package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInt1(t *testing.T) {
	tests := []struct {
		str      string
		expected int64
	}{
		{"-128", 128},
		{"0", 0},
		{"127", 127},
	}
	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.str, 10, 8)
			require.NoError(t, err)
			require.EqualValues(t, tc.expected, got)
		})
	}
}
